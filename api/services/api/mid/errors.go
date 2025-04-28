package mid

import (
	"context"
	"errors"
	"github.com/matheusgcoppi/service/app/sdk/errs"
	"github.com/matheusgcoppi/service/app/sdk/mid"
	"github.com/matheusgcoppi/service/foundation/logger"
	"github.com/matheusgcoppi/service/foundation/web"
	"net/http"
)

// Errors executed the errors middleware functionality.
func Errors(log *logger.Logger) web.MidHandler {
	m := func(next web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return next(ctx, w, r)
			}

			if err := mid.Errors(ctx, log, hdl); err != nil {
				var currentError errs.Error
				errors.As(err, &currentError)
				if err := web.Respond(ctx, w, currentError, currentError.Code.HttpStatus()); err != nil {
					return err
				}

				// If we receive the shutdown err we need to return it
				// to the base handler to shut down the service.
				if web.IsShutdown(err) {
					return err
				}
			}
			return nil
		}

		return h
	}
	return m
}
