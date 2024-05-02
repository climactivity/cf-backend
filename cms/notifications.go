package main

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/cron"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/pocketbase/pocketbase/tools/template"
	"log"
	"net/mail"
	"strconv"
	"time"
)

func sendEmailNotificationToUser(user *models.Record, emailTemplate EmailTemplate, maild mailer.Mailer, app *pocketbase.PocketBase) error {

	html, err := emailTemplate.htmlTemplateRenderer.Render(map[string]any{"AppUrl": app.Settings().Meta.AppUrl})
	plain, err := emailTemplate.plainTextTemplateRenderer.Render(map[string]any{"AppUrl": app.Settings().Meta.AppUrl})

	if err != nil {
		log.Fatal(err)
	}

	message := &mailer.Message{
		From: mail.Address{
			Address: app.Settings().Meta.SenderAddress,
			Name:    app.Settings().Meta.SenderName,
		},
		To:      []mail.Address{{Address: user.Email()}},
		Subject: emailTemplate.subjectLine,
		HTML:    html,
		Text:    plain,
	}

	log.Default().Println("Sending ", message.From, message.To, message.Subject)

	return maild.Send(message)
}

func sendPushNotificationToUser(user *models.Record, app *pocketbase.PocketBase) error {
	//TODO
	log.Default().Println("Would have send a push notification to", user.Email())
	return nil
}

func notifyPushSubscribers(app *pocketbase.PocketBase) error {
	records, err := app.Dao().FindRecordsByExpr(
		"users",
		dbx.HashExp{"notification_push": true},
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		sendPushNotificationToUser(record, app)
	}

	return nil
}

type EmailTemplate struct {
	subjectLine               string
	htmlTemplateRenderer      *template.Renderer
	plainTextTemplateRenderer *template.Renderer
}

func NewEmailTemplate(subjectLine string, htmlViewPath string, plainViewPath string, registry *template.Registry) EmailTemplate {
	return EmailTemplate{
		subjectLine:               subjectLine,
		htmlTemplateRenderer:      registry.LoadFiles(htmlViewPath),
		plainTextTemplateRenderer: registry.LoadFiles(plainViewPath),
	}
}

func getRecipientChunk(app *pocketbase.PocketBase, year int, week int) []*models.Record {

	// todo paginate this to stay within api limits of greensta when we go live
	records, err := app.Dao().FindRecordsByFilter(
		CollectionWeeks,
		"year = {:year} && weeknumber = {:week}",
		"-weeknumber",
		-1,
		0,
		dbx.Params{"year": strconv.Itoa(year), "week": strconv.Itoa(week)},
	)

	if err != nil {
		log.Default().Println(err)
		return nil
	}

	return records
}

func getNextRecipientChunk(app *pocketbase.PocketBase) []*models.Record {
	year, week := GetCurrentYearWeek()

	return getRecipientChunk(app, year, week)
}

func getNextRecipientChunkLastWeek(app *pocketbase.PocketBase) []*models.Record {
	year, week := GetCurrentYearWeek()

	return getRecipientChunk(app, year, week-1)

}

func safeSendReminder(weekRecord *models.Record, hasBeenProcessedKey string, t EmailTemplate, app *pocketbase.PocketBase, maild mailer.Mailer) bool {
	// notification has already been sent
	reminderSent := weekRecord.GetBool(hasBeenProcessedKey)
	if reminderSent {
		return false
	}
	// fetch the user record to check if email notifications are turned on (and to get the email address, ofc)
	userId := weekRecord.GetString("owner")
	if userId == "" {
		log.Default().Println("User was NIL, something's bugged!")
		return false
	}
	userRecord, err := app.Dao().FindRecordById(CollectionUsers, userId)
	if err != nil {
		log.Default().Println(err)
		return false
	}

	if !userRecord.GetBool("notification_email") {
		return false
	}
	// send the email, finally
	if err := sendEmailNotificationToUser(userRecord, t, maild, app); err != nil {
		ReportError(WARN, err)
		return false
	}

	//save information that we sent the message to prevent duplicates
	weekRecord.Set(hasBeenProcessedKey, true)
	if err := app.Dao().SaveRecord(weekRecord); err != nil {
		ReportError(WARN, err)
		return false
	}

	return true
}

func EnqueueSaturdayReminderMessages(app *pocketbase.PocketBase, maild mailer.Mailer, registry *template.Registry) {

	chunk := getNextRecipientChunk(app)

	template := NewEmailTemplate(
		"Aktion GUTES KLIMA Erinnerung",
		"views/HTMLEmailSaturdayReminderTemplate.html",
		"views/PlainEmailSaturdayReminderTemplate.txt",
		registry,
	)

	for _, r := range chunk {
		if r == nil {
			continue
		}
		if r.GetString("state") != string(Completed) {
			safeSendReminder(r, "SaturdayReminderSent", template, app, maild)
		}
	}

}

func EnqueueThursdayReminderMessages(app *pocketbase.PocketBase, maild mailer.Mailer, registry *template.Registry) {

	chunk := getNextRecipientChunk(app)

	template := NewEmailTemplate(
		"Aktion GUTES KLIMA Erinnerung",
		"views/HTMLEmailReminderTemplate.html",
		"views/PlainEmailReminderTemplate.txt",
		registry,
	)

	for _, r := range chunk {
		if r == nil {
			continue
		}
		safeSendReminder(r, "FridayReminderSent", template, app, maild)
	}

	//notifyEmailSubscribers(app)
	notifyPushSubscribers(app)
}

func SetFriday(app *pocketbase.PocketBase) {
	chunk := getNextRecipientChunk(app)
	for _, r := range chunk {

		if r == nil {
			continue
		}
		r.Set("state", string(Pending))
	}
}

func SetMissed(app *pocketbase.PocketBase) {
	chunk := getNextRecipientChunkLastWeek(app)
	for _, r := range chunk {

		if r == nil {
			continue
		}
		if r.Get("completion") == nil {
			r.Set("state", string(Missed))
		}
	}
}

func HandleCronTasks(app *pocketbase.PocketBase) func(e *core.ServeEvent) error {
	return func(e *core.ServeEvent) error {
		scheduler := cron.New()

		location, err := time.LoadLocation("Local")
		if err != nil {
			panic(err)

		}
		scheduler.SetTimezone(location)

		maild := app.NewMailClient()
		registry := template.NewRegistry()

		// see https://crontab.guru/#0_18_*_*_4
		// “At 18:00 on Thursday.”
		scheduler.MustAdd("thursdayReminder", "0 18 * * 4", func() {
			EnqueueThursdayReminderMessages(app, maild, registry)
		})

		scheduler.MustAdd("fridayUpdate", "0 0 * * 5", func() {
			SetFriday(app)
		})

		scheduler.MustAdd("saturdayReminder", "0 18 * * 6", func() {
			EnqueueSaturdayReminderMessages(app, maild, registry)
		})

		scheduler.MustAdd("evalWeek", "59 23 * * 0", func() {
			SetMissed(app)
		})

		scheduler.Start()
		return nil
	}
}
