package mid

import (
	"context"
	"github.com/matheusgcoppi/service/foundation/logger"
	"github.com/matheusgcoppi/service/foundation/web"
	"net/http"
)

func Logger(logger *logger.Logger) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			logger.Info(ctx, "request started", "method", r.Method, r.URL.Path, "remoteaddr", r.RemoteAddr)

			err := handler(ctx, w, r)

			logger.Info(ctx, "request completed", "method", r.Method, r.URL.Path, "remoteaddr", r.RemoteAddr)

			return err
		}

		return h
	}
	return m
}
