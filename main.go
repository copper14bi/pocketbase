package main

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// main is the entry point for the PocketBase fork application.
// It initializes the PocketBase instance and registers any custom
// hooks, routes, or middleware before starting the server.
func main() {
	app := pocketbase.New()

	// Register custom hooks before the app starts
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Register custom API routes
		e.Router.GET("/api/custom/health", func(c echo.Context) error {
			return c.JSON(200, map[string]string{
				"status":  "ok",
				"version": app.Settings().Meta.AppName,
			})
		}, apis.ActivityLogger(app))

		return nil
	})

	// Log record creation events for auditing
	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		app.Logger().Info(
			"Record created",
			"collection", e.Record.Collection().Name,
			"id", e.Record.Id,
		)
		return nil
	})

	// Log record deletion events for auditing
	app.OnRecordAfterDeleteRequest().Add(func(e *core.RecordDeleteEvent) error {
		app.Logger().Info(
			"Record deleted",
			"collection", e.Record.Collection().Name,
			"id", e.Record.Id,
		)
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
