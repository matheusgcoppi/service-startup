package checkapi

import "net/http"

// Routes add specific routes for this group.
func Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /liveness", liveness)
	mux.HandleFunc("GET /readiness", readiness)
}
