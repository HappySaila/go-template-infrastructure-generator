package handlers

import (
	"dev/[[.Dashed]].git/business/[[.Dashed]]"
	"dev/[[.Dashed]].git/business/i"
	"dev/[[.Dashed]].git/business/mid"
	"dev/[[.Dashed]].git/foundation/web"
	"net/http"
	"os"
)

type [[.CapsCamel]] struct {
	Service *[[.Snake]].Service
}

// API constructs a http.Handler with all application routes defined
func API(log i.Logger, [[.Letter]] [[.CapsCamel]], shutdown chan os.Signal) *web.App {

	// Create web app with middleware
	app := web.NewApp(
		shutdown,
		mid.Logger(log),
		mid.Errors(log),
		mid.Panics(log),
	)

	// Check Service
	ch := check{}
	app.Handle(http.MethodGet, "/readiness", ch.readiness)
	app.Handle(http.MethodGet, "/liveliness", ch.liveliness)

	// [[.CapsCamel]] Handlers
	app.Handle(http.MethodPost, "/create", [[.Letter]].create)
	return app

}

// Init will initialise the Service
func Init(db [[.Snake]].Store, log i.Logger) [[.CapsCamel]] {

	// Initialise services
	[[.Letter]] := [[.CapsCamel]]{
		Service: &[[.Snake]].Service{
			Log:   log,
			Store: db,
		},
	}
	return [[.Letter]]

}
