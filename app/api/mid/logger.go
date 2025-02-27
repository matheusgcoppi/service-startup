package mid

import (
	"context"
	"github.com/matheusgcoppi/service/foundation/web"
	"net/http"
)

func Logger(handler web.Handler) web.Handler {
	h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

		// LOGGING HERE

		err := handler(ctx, w, r)

		// LOGGING HERE

		return err
	}

	return h
}
