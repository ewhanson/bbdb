package main

import (
	"github.com/ewhanson/bbdb/hooks"
	"github.com/pocketbase/pocketbase"
	"log"
)

func main() {
	log.Print("Starting BBDB...")
	app := pocketbase.New()

	hooks.Init(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
