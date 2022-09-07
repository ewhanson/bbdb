package notifications

import (
	"context"
	appConfig "github.com/ewhanson/bbdb/config"
	"github.com/ewhanson/bbdb/mailer"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/procyon-projects/chrono"
	"log"
)

// ScheduledNotifications defines group of methods for tracking and sending
// grouped email notifications on updated content
type ScheduledNotifications struct {
	taskScheduler chrono.TaskScheduler
	shouldNotify  bool
	chronInterval string
	config        *appConfig.Config
}

// New runs initial setup for ScheduledNotifications and set up scheduled tasks
func New(app *pocketbase.PocketBase, config *appConfig.Config) *ScheduledNotifications {
	sns := &ScheduledNotifications{
		chronInterval: config.ChronInterval,
		taskScheduler: chrono.NewDefaultTaskScheduler(),
		config:        config,
	}

	_, err := sns.taskScheduler.ScheduleWithCron(func(ctx context.Context) {
		if sns.shouldNotify {
			err := sns.dispatchNotifications(app)
			if err != nil {
				log.Print(err)
				return
			}
		}
	}, sns.chronInterval, chrono.WithLocation("America/Vancouver"))

	if err != nil {
		log.Fatal("Could not setup task scheduler:", err)
	}

	return sns
}

// SetUpdateAvailable flags that the next scheduled check should trigger notification emails
func (sns *ScheduledNotifications) SetUpdateAvailable() {
	sns.shouldNotify = true
}

// dispatchNotifications gets a list of all subscribed email addresses and dispatches a notification email
func (sns *ScheduledNotifications) dispatchNotifications(app *pocketbase.PocketBase) error {
	collection, err := app.Dao().FindCollectionByNameOrId("subscribers")
	if err != nil {
		return err
	}

	var rows []dbx.NullStringMap
	err = app.Dao().RecordQuery(collection).All(&rows)
	if err != nil {
		return err
	}
	records := models.NewRecordsFromNullStringMaps(collection, rows)
	for _, record := range records {
		_ = sns.sendUpdateEmail(record.GetStringDataValue("email"), record.GetId())
	}

	sns.shouldNotify = false

	return nil
}

func (sns *ScheduledNotifications) SendWelcomeEmail(emailAddress string, userId string) error {
	m, err := mailer.New([]string{emailAddress}, "📫 Welcome to Babygramz updates")
	if err != nil {
		return err
	}

	m.AddPlainTextMsg([]string{
		"Welcome! 👋",
		"",
		"You are receiving this because you signed up to receive notification emails from Babygramz.",
		"You will receive an email update, once a day, whenever new photos are available",
		"",
		"---",
		"",
		getPlainTextFooter(userId),
	})
	m.AddHtmlMsg([]string{
		"<html><body style=\"font-family:-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;\">",
		"<h1>Welcome! 👋</h1><p>You are receiving this because you signed up to receive notification emails from <a href=\"https://babygramz.com\">Babygramz</a>.</p><p>You will receive an email update, once per day, whenever new photos are available</p>",
		getHtmlFooter(userId),
		"</body></html>",
	})

	if err = m.Send(); err != nil {
		return err
	}

	return nil
}

// sendUpdateEmail fires off email via SMTP
func (sns *ScheduledNotifications) sendUpdateEmail(emailAddress string, userId string) error {
	m, err := mailer.New([]string{emailAddress}, "📸 New photos available!")
	if err != nil {
		return err
	}

	m.AddPlainTextMsg([]string{
		"Good news! 🎉",
		"",
		"New photos are available on Babygramz. Visit https://babygramz.com to view them.",
		"",
		"---",
		"",
		getPlainTextFooter(userId),
	})
	m.AddHtmlMsg([]string{
		"<html><body style=\"font-family:-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;\">",
		"<h1>Good news! 🎉</h1> <p>New photos are available on Babygramz.</p><p>Visit <a href=\"https://babygramz.com\">babygramz.com</a> to view them.</p>",
		getHtmlFooter(userId),
		"</body></html>",
	})

	if err = m.Send(); err != nil {
		return err
	}

	return nil
}

func getHtmlFooter(userId string) string {
	return "<p style=\"color: grey; font-size: 12.8px; margin: 4px 0 4px 0;\">You are receiving this email because you subscribed to update notifications from babygramz.com.</p><p style=\"color: grey; font-size: 12.8px; margin: 4px 0 4px 0;\">No longer want to receive these updates? <a style=\"color: grey;\" href=\"https://babygramz.com/unsubscribe?id=" + userId + "\">Unsubscribe from these notifications.</a></p>"
}

func getPlainTextFooter(userId string) string {
	return "You are receiving this email because you subscribed to update notifications from babygramz.com.\r\n\r\nNo longer want to receive these updates? Visit the following URL to unsubscribe: https://babygramz.com/unsubscribe?id=" + userId
}
