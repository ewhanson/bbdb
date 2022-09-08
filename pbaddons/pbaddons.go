package pbaddons

import (
	appConfig "github.com/ewhanson/bbdb/config"
	"github.com/ewhanson/bbdb/notifications"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"log"
	"net/http"
	"os"
)

// Init adds custom logic to PocketBase app
func Init(app *pocketbase.PocketBase) {
	sns := initNotifications(app)
	setupImageHeaders(app)
	setupNewPhotoNotifications(app, sns)
	setupSubscribeRecordAction(app, sns)
	addRoutes(app)
}

func addRoutes(app *pocketbase.PocketBase) {
	setupSubscriptionRoutes(app)
	setupApiVersionRoute(app)
}

func initNotifications(app *pocketbase.PocketBase) *notifications.ScheduledNotifications {
	configDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	config, err := appConfig.LoadConfig(configDir)
	if err != nil {
		log.Fatal("Cannot load app config:", err)
	}

	// Add "photos added" email notifications
	return notifications.New(app, &config)
}

func setupImageHeaders(app *pocketbase.PocketBase) {
	// Add cache control headers for image caching
	app.OnFileDownloadRequest().Add(func(e *core.FileDownloadEvent) error {
		// TODO: Specify that cache headers should only apply to images
		e.HttpContext.Response().Header().Set("Cache-Control", "public, max-age=31536000")
		return nil
	})
}

func setupNewPhotoNotifications(app *pocketbase.PocketBase, sns *notifications.ScheduledNotifications) {
	// TODO: Remove before prod
	//sns.SetUpdateAvailable()

	app.OnModelAfterCreate().Add(func(e *core.ModelEvent) error {
		if e.Model.TableName() == "photos" {
			sns.SetUpdateAvailable()
		}

		return nil
	})
}

func setupSubscribeRecordAction(app *pocketbase.PocketBase, sns *notifications.ScheduledNotifications) {
	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		if e.Record.Collection().Name == "subscribers" {
			email := e.Record.GetStringDataValue("email")
			err := sns.SendWelcomeEmail(email, e.Record.GetId())
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func setupSubscriptionRoutes(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		_, err := e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/api/unsubscribe",
			Handler: func(c echo.Context) error {
				id := c.QueryParam("id")

				collection, err := app.Dao().FindCollectionByNameOrId("subscribers")
				if err != nil {
					return err
				}

				record, err := app.Dao().FindRecordById(collection, id, nil)
				if err != nil {
					return err
				}
				err = app.Dao().DeleteRecord(record)
				if err != nil {
					return err
				}

				return c.String(http.StatusOK, "Successfully unsubscribed")
			},
			Middlewares: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

func setupApiVersionRoute(app *pocketbase.PocketBase) {
	configDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	config, err := appConfig.LoadConfig(configDir)
	if err != nil {
		log.Fatal("Cannot load app config:", err)
	}

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		_, err := e.Router.AddRoute(echo.Route{
			Method: "GET",
			Path:   "/api/version",
			Handler: func(c echo.Context) error {
				return c.String(200, config.APIVersion)
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
