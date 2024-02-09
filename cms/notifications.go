package main

import (
	"log"
	"net/mail"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/cron"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/pocketbase/pocketbase/tools/template"
)

func sendEmailNotificationToUser(user *models.Record, template *template.Renderer, app *pocketbase.PocketBase) error {

	html, err := template.Render(map[string]any{"AppUrl": app.Settings().Meta.AppUrl})

	if err != nil {
		log.Fatal(err)
	}

	message := &mailer.Message{
		From: mail.Address{
			Address: app.Settings().Meta.SenderAddress,
			Name:    app.Settings().Meta.SenderName,
		},
		To:      []mail.Address{{Address: user.Email()}},
		Subject: "#ClimateFriday Erinnerung",
		HTML:    html,
		Text: `Hi, 

		  Auch diese Woche machen wieder viele Menschen beim #ClimateFriday> mit.
		  Bist Du diese Woche auch dabei?

		  Viele Grüße,
		  das #ClimateFriday-Team
		`,
	}

	log.Default().Println("Sending ", message.From, message.To, message.Subject)

	return app.NewMailClient().Send(message)
}

var htmlReminderMail = template.NewRegistry().LoadFiles("views/HTMLEmailReminderTemplate.html")

func notifyEmailSubscribers(app *pocketbase.PocketBase) error {
	records, err := app.Dao().FindRecordsByExpr(
		"users",
		dbx.HashExp{"notification_email": true},
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		if err := sendEmailNotificationToUser(record, htmlReminderMail, app); err != nil {
			ReportError(WARN, err)
		}
	}

	return nil
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

func handleNotifications(app *pocketbase.PocketBase) func(e *core.ServeEvent) error {
	return func(e *core.ServeEvent) error {
		scheduler := cron.New()

		// see https://crontab.guru/#0_18_*_*_4
		// “At 18:00 on Thursday.”
		scheduler.MustAdd("hello", "0 18 * * 4", func() {
			notifyEmailSubscribers(app)
			notifyPushSubscribers(app)
		})

		scheduler.Start()
		return nil
	}
}
