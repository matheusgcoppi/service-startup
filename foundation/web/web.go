// Package web contains a small web framework extension.
package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
	"syscall"
	"time"
)

// A Handler is a type that handles a http request within our own little framework
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*http.ServeMux
	shutdown chan os.Signal
	mw       []MidHandler
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal, mw ...MidHandler) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
		mw:       mw,
	}
}

// SignalShutdown is used to gracefully shut down the app when an integrity
// issue is identified
func (app *App) SignalShutdown() {
	app.shutdown <- syscall.SIGTERM
}

// HandleFuncNoMiddleware Handle sets a handler function for a given HTTP method and a path pair
// to the application server mux with no middleware.
func (app *App) HandleFuncNoMiddleware(pattern string, handler Handler, mw ...MidHandler) {
	h := func(w http.ResponseWriter, r *http.Request) {
		v := Values{
			TraceId: uuid.NewString(),
			Now:     time.Now().UTC(),
		}
		ctx := setValues(r.Context(), &v)

		if err := handler(ctx, w, r); err != nil {
			// ERROR HANDLING HERE
			if validateError(err) {
				app.SignalShutdown()
				return
			}
			fmt.Println(err)
			return
		}
	}

	app.ServeMux.HandleFunc(pattern, h)
}

// HandleFunc Handle sets a handler function for a given HTTP method and a path pair
// to the application server mid.
func (app *App) HandleFunc(pattern string, handler Handler, mw ...MidHandler) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(app.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		v := Values{
			TraceId: uuid.NewString(),
			Now:     time.Now().UTC(),
		}
		ctx := setValues(r.Context(), &v)

		if err := handler(ctx, w, r); err != nil {
			// ERROR HANDLING HERE
			if validateError(err) {
				app.SignalShutdown()
				return
			}
			fmt.Println(err)
			return
		}
	}

	app.ServeMux.HandleFunc(pattern, h)
}

// validateError validates the error for special conditions that do not
// warrant an actual shutdown by the system.
func validateError(err error) bool {
	// Ignore syscall.EPIPE and syscall.ECONNRESET errors which occurs
	// when a write operation happens on the http.responseWriter that
	// has a simultaneously been disconnected by the client (TCP connections
	// are broken).For instance, when a large amount of data is being written or streamed to the client.

	switch {
	case errors.Is(err, syscall.EPIPE):
		return false

	case errors.Is(err, syscall.ECONNRESET):
		return false
	}

	return true
}
