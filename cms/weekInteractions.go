package main

import (
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type WeekState string

const (
	DontRemind WeekState = "dont_remind"
	Remind     WeekState = "remind"
	Pending    WeekState = "pending"
	Missed     WeekState = "missed"
	Completed  WeekState = "completed"
)

const (
	CollectionWeeks          string = "weeks"
	CollectionParticipations string = "participations"
	CollectionUsers          string = "users"
)

type WeekInfo struct {
	weekNumber   int
	state        WeekState
	completionId string
}

// GetCurrentYearWeek gets the current year and week (that is really all it does)
func GetCurrentYearWeek() (int, int) {
	now := time.Now().UTC()
	year, week := now.ISOWeek()
	return year, week
}

func getCurrentWeekRecord(app *pocketbase.PocketBase, user *models.Record) (*models.Record, error) {
	year, week := GetCurrentYearWeek()
	return app.Dao().FindFirstRecordByFilter(
		CollectionWeeks,
		"year = {:year} && weeknumber = {:week} && owner = {:user}",
		dbx.Params{"year": strconv.Itoa(year), "week": strconv.Itoa(week), "user": user.Id},
	)

}

// gets the currently active Week
func handleGetCurrentWeek(app *pocketbase.PocketBase) func(c echo.Context) error {
	return func(c echo.Context) error {
		info := apis.RequestInfo(c)
		user, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

		if user == nil {
			return apis.NewUnauthorizedError("Not authorized", "")
		}

		record, err := getCurrentWeekRecord(app, user)
		if err != nil {
			return c.JSON(http.StatusNotFound, record)
		}

		canAccess, err := app.Dao().CanAccessRecord(record, info, record.Collection().ViewRule)

		if err != nil {
			log.Default().Println(err)
			return apis.NewApiError(http.StatusInternalServerError, "Something went wrong :(", nil)
		}

		if !canAccess {
			log.Default().Println(canAccess)

			//return apis.NewForbiddenError("Not allowed", nil)
		}

		return c.JSON(http.StatusOK, record)
	}
}

func initWeek(app *pocketbase.PocketBase, user *models.Record, remind bool) (*models.Record, error) {
	year, week := GetCurrentYearWeek()

	record, _ := getCurrentWeekRecord(app, user)
	if record != nil {
		return record, nil
	}

	collection, err := app.Dao().FindCollectionByNameOrId(CollectionWeeks)
	if err != nil {
		return nil, err
	}

	record = models.NewRecord(collection)
	record.Set("year", year)
	record.Set("weeknumber", week)
	record.Set("owner", user.Id)

	now := time.Now().UTC()

	day := now.Weekday()

	if day == time.Friday {
		record.Set("state", string(Pending))

	} else {
		if remind {
			record.Set("state", string(Remind))
		} else {
			record.Set("state", string(DontRemind))
		}

	}

	if err := app.Dao().SaveRecord(record); err != nil {
		return nil, err
	}

	return record, nil

}

func handleInitWeek(app *pocketbase.PocketBase) func(c echo.Context) error {
	return func(c echo.Context) error {

		user, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

		if user == nil {
			return apis.NewUnauthorizedError("Not authorized", "")
		}

		data := apis.RequestInfo(c).Data
		remind := data["remind"]

		// get current week if exists and updated it with new reminder setting
		record, _ := getCurrentWeekRecord(app, user)
		if record != nil {

			newRemind := remind != nil && remind != ""

			// update reminder settings when user tries to recreate an already existing record. Something has gone
			// slightly wrong, but the intention is still clear
			if newRemind != (record.GetString("state") == string(Remind)) {
				if newRemind {
					record.Set("state", string(Remind))
				} else {
					record.Set("state", string(DontRemind))
				}

				if err := app.Dao().SaveRecord(record); err != nil {
					return err
				}
				return c.JSON(http.StatusOK, record)

			}

			return c.JSON(http.StatusNotModified, record)
		}

		// create new record
		record, err := initWeek(app, user, remind != nil && remind != "")

		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "error", err)
		}

		return c.JSON(http.StatusCreated, record)
	}
}

func handleCompleteWeek(app *pocketbase.PocketBase) func(c echo.Context) error {
	return func(c echo.Context) error {

		// read/scan the request body fields into a typed struct
		// (note that a body cannot be read twice with Bind because it is a stream)
		data := struct {
			Option string `json:"option" form:"option"`
			Note   string `json:"note" form:"note"`
		}{}
		if err := c.Bind(&data); err != nil {
			return apis.NewBadRequestError("Failed to read request data", err)
		}

		user, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)
		weekRecord, _ := getCurrentWeekRecord(app, user)
		if weekRecord == nil {
			// user started week on friday and hasn't created an entry beforehand
			newRecord, err := initWeek(app, user, false)
			if err != nil {
				return apis.NewApiError(http.StatusInternalServerError, "error", err)
			}
			//assign new weekRecord
			weekRecord = newRecord
		}

		participationsCollection, err := app.Dao().FindCollectionByNameOrId(CollectionParticipations)
		if err != nil {
			return apis.NewBadRequestError("Something went wrong", err)
		}

		year, week := GetCurrentYearWeek()

		participationRecord := models.NewRecord(participationsCollection)

		participationRecord.Set("owner", user.Id)
		participationRecord.Set("option", data.Option)
		participationRecord.Set("note", data.Note)
		participationRecord.Set("year", year)
		participationRecord.Set("weeknumber", week)

		if err := app.Dao().SaveRecord(participationRecord); err != nil {
			log.Default().Println(fmt.Errorf("failed to save record: %v", err))

			return apis.NewBadRequestError("Failed to read request data", err)
		}

		if errs := app.Dao().ExpandRecord(weekRecord, []string{"completion"}, nil); len(errs) > 0 {
			log.Default().Println(fmt.Errorf("failed to expand: %v", errs))
			return apis.NewBadRequestError("Failed to read request data", err)
		}

		return c.JSON(http.StatusCreated, weekRecord)
	}
}

func handleGetCompleteWeek(app *pocketbase.PocketBase) func(c echo.Context) error {
	return func(c echo.Context) error {

		user, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)
		weekRecord, err := getCurrentWeekRecord(app, user)

		if errs := app.Dao().ExpandRecord(weekRecord, []string{"completion"}, nil); len(errs) > 0 {
			log.Default().Println(fmt.Errorf("failed to expand: %v", errs))
			return apis.NewBadRequestError("Failed to read request data", err)
		}

		return c.JSON(http.StatusCreated, weekRecord)
	}

}

func CurrentWeekHandler(app *pocketbase.PocketBase) func(e *core.ServeEvent) error {
	return func(e *core.ServeEvent) error {

		// init week -> "Ich bin diese Woche dabei!"
		e.Router.GET("/api/currentweek", handleGetCurrentWeek(app))
		e.Router.POST("/api/currentweek", handleInitWeek(app))

		// complete week -> participated on a friday
		e.Router.GET("/api/currentweek/completion", handleGetCompleteWeek(app))
		e.Router.POST("/api/currentweek/completion", handleCompleteWeek(app))

		return nil
	}
}

func DebugHandler(app *pocketbase.PocketBase) func(e *core.ServeEvent) error {
	return func(e *core.ServeEvent) error {
		maild := app.NewMailClient()
		registry := template.NewRegistry()

		e.Router.POST("/api/debug/sentThursdayMail", func(c echo.Context) error {
			admin, _ := c.Get(apis.ContextAdminKey).(*models.Admin)
			if admin == nil {
				return apis.NewForbiddenError("only admins can do this", nil)
			}
			EnqueueThursdayReminderMessages(app, maild, registry)
			return c.NoContent(http.StatusOK)
		})

		e.Router.POST("/api/debug/sentSaturdayMail", func(c echo.Context) error {
			admin, _ := c.Get(apis.ContextAdminKey).(*models.Admin)
			if admin == nil {
				return apis.NewForbiddenError("only admins can do this", nil)
			}
			EnqueueSaturdayReminderMessages(app, maild, registry)
			return c.NoContent(http.StatusOK)
		})

		e.Router.POST("/api/debug/setFriday", func(c echo.Context) error {
			admin, _ := c.Get(apis.ContextAdminKey).(*models.Admin)
			if admin == nil {
				return apis.NewForbiddenError("only admins can do this", nil)
			}
			SetFriday(app)
			return c.NoContent(http.StatusOK)
		})

		e.Router.POST("/api/debug/setMissed", func(c echo.Context) error {
			admin, _ := c.Get(apis.ContextAdminKey).(*models.Admin)
			if admin == nil {
				return apis.NewForbiddenError("only admins can do this", nil)
			}
			SetMissed(app)
			return c.NoContent(http.StatusOK)
		})

		return nil
	}
}
