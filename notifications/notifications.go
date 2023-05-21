package notifications

import (
	"github.com/ewhanson/bbdb/photos_queue"
	"github.com/ewhanson/bbdb/scheduler"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/spf13/viper"
	"log"
	"net/mail"
	"strconv"
)

// ScheduledNotifications defines group of methods for tracking and sending
// grouped email notifications on updated content
type ScheduledNotifications struct {
	app       *pocketbase.PocketBase
	scheduler *scheduler.Scheduler
	pq        *photos_queue.PhotosQueue
}

// New runs initial setup for ScheduledNotifications and set up scheduled tasks
func New(app *pocketbase.PocketBase, s *scheduler.Scheduler, pq *photos_queue.PhotosQueue) *ScheduledNotifications {
	sns := &ScheduledNotifications{
		app:       app,
		scheduler: s,
		pq:        pq,
	}

	notificationTime := viper.GetString("notificationTime")
	_, err := sns.scheduler.GetScheduler().Every(1).Day().At(notificationTime).Do(func() {
		if sns.shouldNotify() {
			err := sns.pq.CleanupNonPendingItems()
			err = sns.dispatchNotifications(app)
			if err != nil {
				log.Println(err.Error())
				return
			}
			err = pq.MarkItemsNotPending()
			if err != nil {
				log.Println(err.Error())
			}
		}
	})
	if err != nil {
		return nil
	}

	return sns
}

func (sns *ScheduledNotifications) shouldNotify() bool {
	return sns.pq.HasPendingItems()

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
	photoCount, _ := sns.pq.GetPendingItemsCount()
	for _, record := range records {
		err = sns.sendUpdateEmail(record.GetString("email"), record.GetString("name"), record.GetId(), photoCount)
		if err != nil {
			failCount++
			log.Println("[Mail error] ", err.Error())
		} else {
			successCount++
		}
	}
	log.Println("[Batch notification dispatch] Succeeded: " + strconv.Itoa(successCount) + ", Failed: " + strconv.Itoa(failCount) + ", Total: " + strconv.Itoa(successCount+failCount))
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
func (sns *ScheduledNotifications) sendUpdateEmail(emailAddress string, name string, userId string, photoCount int) error {
	m := sns.app.NewMailClient()

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
