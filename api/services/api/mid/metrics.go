package mid

import (
	"context"
	"github.com/matheusgcoppi/service/app/sdk/mid"
	"github.com/matheusgcoppi/service/foundation/web"
	"net/http"
)

// Metrics updates program counters using middleware functionality.
func Metrics() web.MidHandler {
	m := func(next web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return next(ctx, w, r)
			}

			return mid.Metrics(ctx, hdl)
		}

		return h
	}

	return m
}
