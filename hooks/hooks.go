package hooks

import (
	"fmt"
	"github.com/disintegration/imaging"
	_ "github.com/ewhanson/bbdb/migrations"
	"github.com/ewhanson/bbdb/notifications"
	"github.com/ewhanson/bbdb/photos_queue"
	"github.com/ewhanson/bbdb/scheduler"
	"github.com/ewhanson/bbdb/ui"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/list"
	"github.com/spf13/viper"
	"io/fs"
	"log"
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
	downloadWebFriendlyImages(app)

	s := scheduler.New()
	pq := photos_queue.New(app, s)

	sns := notifications.New(app, s, pq)
	setupSubscribeRecordAction(app, sns)

	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		s.Start()
		return nil
	})
}

func extendRootCmd(app *pocketbase.PocketBase) {
	// Also add in default migration command
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true,
		Dir:         "migrations",
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
}

// bindStaticFrontendUI registers the endpoints that serve the static frontend UI.
func bindStaticFrontendUI(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Serves static files from the ui/dist directory
		e.Router.GET("/*", staticDirectoryHandler(ui.DistDirFS, false), middleware.Gzip())

		return nil
	})
}

func setupSubscribeRecordAction(app *pocketbase.PocketBase, sns *notifications.ScheduledNotifications) {
	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		if e.Record.Collection().Name == "subscribers" {
			email := e.Record.GetString("email")
			name := e.Record.GetString("name")
			err := sns.SendWelcomeEmail(email, name, e.Record.GetId())
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

// staticDirectoryHandler is similar to `apis.StaticDirectoryHandler`
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

func downloadWebFriendlyImages(app *pocketbase.PocketBase) {
	var imageContentTypes = []string{"image/png", "image/jpg", "image/jpeg", "image/gif"}

	app.OnFileDownloadRequest().Add(func(e *core.FileDownloadEvent) error {
		if e.HttpContext.QueryParam("size") != "" {
			filesystem, err := app.NewFilesystem()
			if err != nil {
				return apis.NewBadRequestError("Filesystem initialization failure.", err)
			}
			defer func() {
				err := filesystem.Close()
				if err != nil {
					log.Println("[Filesystem error]: ", err)
				}
			}()

			filename := e.HttpContext.PathParam("filename")
			originalPath := e.Record.BaseFilesPath() + "/" + filename

			// Extract the original file meta attributes and check its existence
			oAttrs, oAttrsErr := filesystem.Attributes(originalPath)
			if oAttrsErr != nil {
				return apis.NewNotFoundError("", err)
			}

			// Check if it's an image
			if list.ExistInSlice(oAttrs.ContentType, imageContentTypes) {
				// Add a thumb size as file suffix
				servedName := "size-large" + "_" + filename
				servedPath := e.Record.BaseFilesPath() + "/thumbs_" + filename + "/" + servedName

				// check if the thumb exists:
				// - if doesn't exist - create a new thumb with the specified thumb size
				// - if exists - compare last modified dates to determine whether the thumb should be recreated
				tAttr, tAttrErr := filesystem.Attributes(servedPath)
				if tAttrErr != nil || oAttrs.ModTime.After(tAttr.ModTime) {
					imgFile, err := filesystem.GetFile(originalPath)
					if err != nil {
						return apis.NewNotFoundError("", err)
					}
					defer func() {
						err := imgFile.Close()
						if err != nil {
							log.Println("[Filesystem error]: ", err)
						}
					}()
					img, err := imaging.Decode(imgFile, imaging.AutoOrientation(true))
					if err != nil {
						return apis.NewNotFoundError("", err)
					}
					thumbSize := ""
					if img.Bounds().Max.X > img.Bounds().Max.Y {
						thumbSize = "1280x0"
					} else {
						thumbSize = "0x1280"
					}
					if err := filesystem.CreateThumb(originalPath, servedPath, thumbSize); err != nil {
						servedPath = originalPath
					}
				}
				e.ServedName = servedName
				e.ServedPath = servedPath
			}
		}
		return nil
	})
}
