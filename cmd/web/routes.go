package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// standard normal middleware
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// a separate middle ware for specific routes the need the session functinality
	dynamicMiddleware := alice.New(app.session.Enable)
	// --------------
	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// secureHdrs := secureHeaders(mux)
	// logReq := app.logRequest(secureHdrs)
	// handler := app.recoverPanic(logReq)

	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	// return handler
	return standardMiddleware.Then(mux)
}
