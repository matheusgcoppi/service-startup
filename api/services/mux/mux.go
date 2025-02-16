// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"github.com/matheusgcoppi/service/api/services/route/sys/checkapi"
	"net/http"
)

// WebApi constructs an http.Handler with all application routes bound.
func WebApi() *http.ServeMux {
	mux := http.NewServeMux()

	checkapi.Routes(mux)

	return mux
}
