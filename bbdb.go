package main

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		subFs := echo.MustSubFS(e.Router.Filesystem, "pb_public")
		e.Router.GET("/*", apis.StaticDirectoryHandler(subFs, false))

		return nil
	})

	app.OnFileDownloadRequest().Add(func(e *core.FileDownloadEvent) error {
		// Add cache control headers for image caching
		// TODO: Specify that cache headers should only apply to images
		e.HttpContext.Response().Header().Set("Cache-Control", "public, max-age=31536000")
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
