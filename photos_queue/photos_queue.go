package photos_queue

import (
	"github.com/ewhanson/bbdb/scheduler"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"log"
)

type PhotosQueue struct {
	app *pocketbase.PocketBase
}

// New initializes PhotoQueue with reference to Pocketbase app and initializes hooks
func New(app *pocketbase.PocketBase, s *scheduler.Scheduler) *PhotosQueue {
	pq := &PhotosQueue{
		app: app,
	}

	pq.addHooks(s)
	return pq
}

// HasPendingItems indicates if there are any records flagged as pending
func (pq *PhotosQueue) HasPendingItems() bool {
	records, err := pq.getPendingItems()
	if err != nil {
		return false
	}

	return len(records) > 0
}

// GetPendingItemsCount indicates how many items are flagged as pending
func (pq *PhotosQueue) GetPendingItemsCount() (int, error) {
	records, err := pq.getPendingItems()
	if err != nil {
		return 0, err
	}

	return len(records), nil
}

// MarkItemsNotPending switches all existing pending items to not pending
func (pq *PhotosQueue) MarkItemsNotPending() error {
	records, err := pq.getPendingItems()
	if err != nil {
		return err
	}

	for _, record := range records {
		record.Set("is_pending", false)
		if err := pq.app.Dao().SaveRecord(record); err != nil {
			log.Println("[Photos queue error] ", err.Error())
		}
	}

	return nil
}

// CleanupNonPendingItems deletes any non-pending items older than week
func (pq *PhotosQueue) CleanupNonPendingItems() error {
	if err := pq.removeNonPendingItems(); err != nil {
		return err
	}
	return nil
}

// addHooks adds functions to necessary pocketbase lifecycle hooks
func (pq *PhotosQueue) addHooks(s *scheduler.Scheduler) {
	pq.addToQueueOnNewPhotoAdded()
	pq.scheduleOutdatedNewPhotosCleanup(s)
}

// getPendingItems queries 'photos_queue' collection for items flagged as pending
func (pq *PhotosQueue) getPendingItems() ([]*models.Record, error) {
	records, err := pq.app.Dao().FindRecordsByExpr("photos_queue", dbx.HashExp{"is_pending": true})
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (pq *PhotosQueue) addToQueueOnNewPhotoAdded() {
	pq.app.OnModelAfterCreate().Add(func(e *core.ModelEvent) error {
		if e.Model.TableName() == "photos" {
			collection, err := pq.app.Dao().FindCollectionByNameOrId("photos_queue")
			if err != nil {
				return err
			}

			record := models.NewRecord(collection)
			record.Set("photo", e.Model.GetId())
			record.Set("is_pending", true)

			if err := pq.app.Dao().SaveRecord(record); err != nil {
				return err
			}
		}

		return nil
	})
}

func (pq *PhotosQueue) scheduleOutdatedNewPhotosCleanup(s *scheduler.Scheduler) {
	_, err := s.GetScheduler().Every(1).Day().Do(pq.performOutdatedNewPhotosCleanup)
	if err != nil {
		log.Println("[Cleanup photos queue scheduling error] ", err.Error())
	}
}

func (pq *PhotosQueue) performOutdatedNewPhotosCleanup() {
	records, err := pq.app.Dao().FindRecordsByExpr("photos_queue",
		dbx.NewExp("created <= datetime('now', '-7 days')"),
	)
	if err != nil && err.Error() != "sql: no rows in result set" {
		log.Println("[Error finding outdated photos queue data] ", err.Error())
		return
	}

	for _, record := range records {
		if err := pq.app.Dao().DeleteRecord(record); err != nil {
			log.Println("[Delete outdated new photo status error] ", err.Error())
		}
	}
}

func (pq *PhotosQueue) removeNonPendingItems() error {
	records, err := pq.app.Dao().FindRecordsByExpr("photos_queue",
		dbx.HashExp{"is_pending": false},
	)
	if err != nil {
		log.Println("[Non-pending cleanup error] ", err.Error())
		return err
	}

	for _, record := range records {
		if err := pq.app.Dao().DeleteRecord(record); err != nil {
			log.Println("[Delete non-pending error] ", err.Error())
		}
	}

	return nil

}
