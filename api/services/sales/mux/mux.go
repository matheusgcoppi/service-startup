// Package mid provides support to bind domain level routes
// to the application mid.
package mux

import (
	"github.com/matheusgcoppi/service/api/services/api/mid"
	"github.com/matheusgcoppi/service/api/services/sales/route/sys/checkapi"
	"github.com/matheusgcoppi/service/foundation/logger"
	"github.com/matheusgcoppi/service/foundation/web"
	"os"
)

// WebApi constructs an http.Handler with all application routes bound.
func WebApi(log *logger.Logger, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log))
	//mid := http.NewServeMux()

	checkapi.Routes(mux)

	return mux
}
