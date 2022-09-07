package main

import (
	"github.com/ewhanson/bbdb/pbaddons"
	"github.com/pocketbase/pocketbase"
	"log"
)

func main() {
	app := pocketbase.New()

	pbaddons.Init(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
