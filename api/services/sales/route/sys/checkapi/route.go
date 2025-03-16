package checkapi

import (
	"github.com/matheusgcoppi/service/foundation/web"
)

// Routes add specific routes for this group.
func Routes(app *web.App) {
	app.HandleFunc("GET /liveness", liveness)
	app.HandleFunc("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testError)
	app.HandleFunc("GET /testpanic", testPanic)
}
