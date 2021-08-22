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
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)
	// --------------

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	// user authetincation routes --------------
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser))

	// static file server ------------------

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// secureHdrs := secureHeaders(mux)
	// logReq := app.logRequest(secureHdrs)
	// handler := app.recoverPanic(logReq)

	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	// return handler
	return standardMiddleware.Then(mux)
}
