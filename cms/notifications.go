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
)

func sendEmailNotificationToUser(user *models.Record, app *pocketbase.PocketBase) error {

	message := &mailer.Message{
		From: mail.Address{
			Address: app.Settings().Meta.SenderAddress,
			Name:    app.Settings().Meta.SenderName,
		},
		To:      []mail.Address{{Address: user.Email()}},
		Subject: "YOUR_SUBJECT...",
		HTML:    "YOUR_HTML_BODY...",
		// bcc, cc, attachments and custom headers are also supported...
	}
	// app.Settings().Meta.
	log.Default().Println("Sending ", message.From, message.To, message.Subject)

	return app.NewMailClient().Send(message)
}

func notifyEmailSubscribers(app *pocketbase.PocketBase) error {
	records, err := app.Dao().FindRecordsByExpr(
		"users",
		dbx.HashExp{"notification_email": true},
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		sendEmailNotificationToUser(record, app)
		if err != nil {
			log.Default().Println(err)
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

		// prints "Hello!" every 2 minutes
		scheduler.MustAdd("hello", "*/2 * * * *", func() {
			notifyEmailSubscribers(app)
			notifyPushSubscribers(app)
		})

		scheduler.Start()
		return nil
	}
}
