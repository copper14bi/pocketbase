// Package hooks provides custom PocketBase hook implementations
// for extending the default behavior of the application.
package hooks

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// RegisterHooks attaches all custom application hooks to the PocketBase instance.
// This is the central place to wire up event listeners and middleware.
func RegisterHooks(app *pocketbase.PocketBase) {
	registerOnBeforeServe(app)
	registerRecordHooks(app)
}

// registerOnBeforeServe sets up hooks that run before the HTTP server starts.
func registerOnBeforeServe(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		log.Println("[hooks] Server is starting up...")

		// Register any custom routes here
		e.Router.GET("/api/health", func(c echo.Context) error {
			return c.JSON(200, map[string]string{
				"status": "ok",
			})
		})

		return nil
	})
}

// registerRecordHooks sets up hooks for record create/update/delete events.
func registerRecordHooks(app *pocketbase.PocketBase) {
	// Log record creation across all collections
	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		log.Printf("[hooks] Record created in collection '%s': %s\n",
			e.Record.Collection().Name,
			e.Record.Id,
		)
		return nil
	})

	// Log record updates across all collections
	app.OnRecordAfterUpdateRequest().Add(func(e *core.RecordUpdateEvent) error {
		log.Printf("[hooks] Record updated in collection '%s': %s\n",
			e.Record.Collection().Name,
			e.Record.Id,
		)
		return nil
	})

	// Log record deletions across all collections
	// NOTE: deletion events do not carry the full record fields, only Id and collection info.
	// This means you cannot access e.Record.Get("someField") here — only Id and Collection().
	app.OnRecordAfterDeleteRequest().Add(func(e *core.RecordDeleteEvent) error {
		log.Printf("[hooks] Record deleted from collection '%s': %s\n",
			e.Record.Collection().Name,
			e.Record.Id,
		)
		return nil
	})
}
