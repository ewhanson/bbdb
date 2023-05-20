package notifications

import (
	"github.com/go-co-op/gocron"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/spf13/viper"
	"log"
	"net/mail"
	"strconv"
	"time"
)

// ScheduledNotifications defines group of methods for tracking and sending
// grouped email notifications on updated content
type ScheduledNotifications struct {
	scheduler    *gocron.Scheduler
	shouldNotify bool
	photoCount   int
	app          *pocketbase.PocketBase
}

// New runs initial setup for ScheduledNotifications and set up scheduled tasks
func New(app *pocketbase.PocketBase) *ScheduledNotifications {
	sns := &ScheduledNotifications{
		scheduler: gocron.NewScheduler(time.Local),
		app:       app,
	}

	notificationTime := viper.GetString("notificationTime")
	_, err := sns.scheduler.Every(1).Day().At(notificationTime).Do(func() {
		if sns.shouldNotify {
			err := sns.dispatchNotifications(app)
			if err != nil {
				log.Print(err)
				return
			}
		}
	})
	if err != nil {
		return nil
	}

	sns.scheduler.StartAsync()

	return sns
}

// SetUpdateAvailable flags that the next scheduled check should trigger notification emails
func (sns *ScheduledNotifications) SetUpdateAvailable() {
	sns.shouldNotify = true
	sns.photoCount += 1
}

// DebugDispatch manually triggers sending of notification emails
func (sns *ScheduledNotifications) DebugDispatch(app *pocketbase.PocketBase) error {
	return sns.dispatchNotifications(app)
}

// GetStatus returns info for dispatch debugging via API
func (sns *ScheduledNotifications) GetStatus() interface{} {
	return struct {
		PhotoCount   int
		ShouldNotify bool
	}{
		PhotoCount:   sns.photoCount,
		ShouldNotify: sns.shouldNotify,
	}
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
	successCount := 0
	failCount := 0
	for _, record := range records {
		err = sns.sendUpdateEmail(record.GetString("email"), record.GetString("name"), record.GetId())
		if err != nil {
			failCount++
			log.Println("[Mail error] ", err.Error())
		} else {
			successCount++
		}
	}
	log.Print("[Batch notification dispatch] Succeeded: " + strconv.Itoa(successCount) + ", Failed: " + strconv.Itoa(failCount) + ", Total: " + strconv.Itoa(successCount+failCount))

	sns.shouldNotify = false
	sns.photoCount = 0

	return nil
}

func (sns *ScheduledNotifications) SendWelcomeEmail(emailAddress string, name string, userId string) error {
	m := sns.app.NewMailClient()

	body := "<html><body style=\"font-family:-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;\"><h1 style=\"display: flex; justify-content: center; margin-bottom: 4rem;\">Babygramz👶🎆</h1><p><strong>Welcome, " + name + "!</strong> 👋</p><p>Thanks for signing up for Babygramz notifications.</p><p>You will receive an email update whenever new photos are posted.</p>" + getHtmlFooter(userId) + "</body></html>"

	err := m.Send(&mailer.Message{
		From: mail.Address{
			Name:    sns.app.Settings().Meta.SenderName,
			Address: sns.app.Settings().Meta.SenderAddress,
		},
		To:      []mail.Address{{Address: emailAddress}},
		Subject: "📫 Welcome to Babygramz",
		HTML:    body,
	})

	if err != nil {
		return err
	}

	return nil
}

// sendUpdateEmail fires off email via SMTP
func (sns *ScheduledNotifications) sendUpdateEmail(emailAddress string, name string, userId string) error {
	m := sns.app.NewMailClient()

	photoCount := sns.photoCount
	photoNoun := "photo"
	if photoCount > 1 {
		photoNoun = "photos"
	}

	body := "<html><body style=\"font-family:-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;\"><h1 style=\"display: flex; justify-content: center; margin-bottom: 4rem;\">Babygramz👶🎆</h1><p><strong>Hey " + name + ",</strong></p><p>Julian has " + strconv.Itoa(photoCount) + " new " + photoNoun + " on Babygramz!</p><p><a href=\"https://babygramz.com/feed\">View on Babygramz.</a></p>" + getHtmlFooter(userId) + "</body></html>"

	err := m.Send(&mailer.Message{
		From: mail.Address{
			Name:    sns.app.Settings().Meta.SenderName,
			Address: sns.app.Settings().Meta.SenderAddress,
		},
		To:      []mail.Address{{Address: emailAddress}},
		Subject: "Update: 📸 " + strconv.Itoa(photoCount) + " new " + photoNoun + " available",
		HTML:    body,
	})

	if err != nil {
		return err
	}

	return nil
}

func getHtmlFooter(userId string) string {
	return "<p style=\"color: grey; font-size: 12.8px; margin: 4px 0 4px 0;\">You are receiving this email because you subscribed to update notifications.</p><p style=\"color: grey; font-size: 12.8px; margin: 4px 0 4px 0;\"><a style=\"color: grey;\" href=\"https://babygramz.com/unsubscribe?id=" + userId + "\">Unsubscribe.</a></p>"
}
