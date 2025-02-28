// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"github.com/matheusgcoppi/service/api/services/route/sys/checkapi"
	"github.com/matheusgcoppi/service/app/api/mid"
	"github.com/matheusgcoppi/service/foundation/logger"
	"github.com/matheusgcoppi/service/foundation/web"
	"os"
)

// WebApi constructs an http.Handler with all application routes bound.
func WebApi(log *logger.Logger, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log))
	//mux := http.NewServeMux()

	checkapi.Routes(mux)

	return mux
}
