package checkapi

import (
	"github.com/matheusgcoppi/service/foundation/web"
)

// Routes add specific routes for this group.
func Routes(app *web.App) {
	app.HandleFuncNoMiddleware("GET /liveness", liveness)
	app.HandleFuncNoMiddleware("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testError)
	app.HandleFunc("GET /testpanic", testPanic)
}
