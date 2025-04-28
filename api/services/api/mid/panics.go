package mid

import (
	"context"
	"github.com/matheusgcoppi/service/app/sdk/mid"
	"github.com/matheusgcoppi/service/foundation/web"
	"net/http"
)

// Panics execute the panic middleware functionality.
func Panics() web.MidHandler {
	m := func(next web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			hdl := func(ctx context.Context) error {
				return next(ctx, w, r)
			}

			return mid.Panics(ctx, hdl)
		}
		return h
	}
	return m
}
