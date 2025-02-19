// Package web contains a small web framework extension.
package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

// A Handler is a type that handles a http request within our own little framework
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*http.ServeMux
	shutdown chan os.Signal
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
	}
}

// HandleFunc Handle sets a handler function for a given HTTP method and a path pair
// to the applciation server mux.
func (app *App) HandleFunc(pattern string, handler Handler) {

	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(r.Context(), w, r); err != nil {
			// ERROR HANDLING HERE
			fmt.Println(err)
			return
		}
	}

	app.ServeMux.HandleFunc(pattern, h)
}
