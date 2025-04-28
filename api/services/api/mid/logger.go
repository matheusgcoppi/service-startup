package mid

import (
	"context"
	"github.com/matheusgcoppi/service/app/sdk/mid"
	"github.com/matheusgcoppi/service/foundation/logger"
	"github.com/matheusgcoppi/service/foundation/web"
	"net/http"
)

func Logger(logger *logger.Logger) web.MidHandler {
	m := func(next web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return next(ctx, w, r)
			}

			return mid.Logger(ctx, logger, r.URL.Path, r.URL.RawQuery, r.Method, r.RemoteAddr, hdl)
		}

		return h
	}
	return m
}
