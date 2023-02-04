package pbaddons

import (
	"fmt"
	"github.com/ewhanson/bbdb/notifications"
	"github.com/ewhanson/bbdb/ui"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/spf13/viper"
	"io/fs"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

var Version string = "(Un-versioned)"

// Init adds custom logic to PocketBase app
func Init(app *pocketbase.PocketBase) {
	extendRootCmd(app)

	addRoutes(app)
	setupImageHeaders(app)
	getPhotoExifDataBeforeCreate(app)

	sns := notifications.New(app)
	setupNewPhotoNotifications(app, sns)
	setupSubscribeRecordAction(app, sns)
}

func extendRootCmd(app *pocketbase.PocketBase) {
	// Also add in default migration command
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})

	app.RootCmd.Version = Version
	app.RootCmd.Use = "bbdb"
	app.RootCmd.Short = "bbdb CLI"

	app.RootCmd.PersistentFlags().StringP(
		"notificationTime",
		"t",
		"HH:mm",
		"Time of day to send notification in HH:mm format")
	_ = viper.BindPFlag("notificationTime", app.RootCmd.PersistentFlags().Lookup("notificationTime"))
	viper.SetDefault("notificationTime", "08:15")
}

func addRoutes(app *pocketbase.PocketBase) {
	setupSubscriptionRoutes(app)
	bindStaticFrontendUI(app)
	setupHealthCheckRoute(app)
}

// bindStaticFrontendUI registers the endpoints that serve the static frontend UI.
func bindStaticFrontendUI(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Serves static files from the ui/dist directory
		e.Router.GET("/*", staticDirectoryHandler(ui.DistDirFS, false), middleware.Gzip())

		return nil
	})
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
			email := e.Record.GetString("email")
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

				record, err := app.Dao().FindRecordById("subscribers", id, nil)
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
			return nil
		}

		dateTaken, err := metaData.DateTime()
		if err != nil {
			return nil
		}

		e.Record.Set("dateTaken", dateTaken.UTC())
		return nil
	})
}

// StaticDirectoryHandler is similar to `apis.StaticDirectoryHandler`
// but will fall back to index.html for SPA routing when returning a 404
//
// see apis.StaticDirectoryHandler for more info on code below
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
	}
}

func setupHealthCheckRoute(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		_, err := e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/healthcheck",
			Handler: func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
			},
			Middlewares: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
