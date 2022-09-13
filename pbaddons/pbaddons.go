package pbaddons

import (
	"fmt"
	appConfig "github.com/ewhanson/bbdb/config"
	"github.com/ewhanson/bbdb/notifications"
	"github.com/ewhanson/bbdb/ui"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/rwcarlsen/goexif/exif"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

// Init adds custom logic to PocketBase app
func Init(app *pocketbase.PocketBase) {
	sns := initNotifications(app)
	setupImageHeaders(app)
	setupNewPhotoNotifications(app, sns)
	setupSubscribeRecordAction(app, sns)
	addRoutes(app)
	getPhotoExifDataBeforeCreate(app)
}

func addRoutes(app *pocketbase.PocketBase) {
	setupSubscriptionRoutes(app)
	setupApiVersionRoute(app)
	bindStaticFrontendUI(app)
}

// bindStaticFrontendUI registers the endpoints that serve the static frontend UI.
func bindStaticFrontendUI(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Serves static files from the ui/dist directory
		e.Router.GET("/*", staticDirectoryHandler(ui.DistDirFS, false), middleware.Gzip())

		return nil
	})
}

func initNotifications(app *pocketbase.PocketBase) *notifications.ScheduledNotifications {
	config, err := appConfig.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load app config: ", err)
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
	config, err := appConfig.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load app config: ", err)
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

func getPhotoExifDataBeforeCreate(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		if e.Record.Collection().Name != "photos" {
			return nil
		}

		file, _, err := e.HttpContext.Request().FormFile("file")
		if err != nil {
			return err
		}

		metaData, err := exif.Decode(file)
		if err != nil {
			return err
		}

		dateTaken, err := metaData.DateTime()
		if err != nil {
			return err
		}

		e.Record.SetDataValue("dateTaken", dateTaken)
		return nil
	})
}

// StaticDirectoryHandler is similar to `apis.StaticDirectoryHandler`
// but will fall back to index.html for SPA routing when returning a 404
//
// @see https://github.com/labstack/echo/issues/2211
func staticDirectoryHandler(fileSystem fs.FS, disablePathUnescaping bool) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.PathParam("*")
		if !disablePathUnescaping { // when router is already unescaping we do not want to do is twice
			tmpPath, err := url.PathUnescape(p)
			if err != nil {
				return fmt.Errorf("failed to unescape path variable: %w", err)
			}
			p = tmpPath
		}

		// fs.FS.Open() already assumes that file names are relative to FS root path and considers name with prefix `/` as invalid
		name := filepath.ToSlash(filepath.Clean(strings.TrimPrefix(p, "/")))

		initialResult := c.FileFS(name, fileSystem)
		if initialResult != nil {
			secondResult := c.FileFS(".", fileSystem)
			return secondResult
		}

		return initialResult

		return initialResult
	}
}
