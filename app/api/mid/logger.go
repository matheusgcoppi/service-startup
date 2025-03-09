package mid

import (
	"context"
	"fmt"
	"github.com/matheusgcoppi/service/foundation/logger"
	"github.com/matheusgcoppi/service/foundation/web"
	"time"
)

// Logger writes information about the request to the logs.
func Logger(ctx context.Context, logger *logger.Logger, path string, rawQuery string, method string, remoteAddr string, handler Handler) error {
	v := web.GetValues(ctx)

	if rawQuery != "" {
		path = fmt.Sprintf("%s?%s", path, rawQuery)
	}

	logger.Info(ctx, "request started", "method", method, "path", path, "remoteaddr", remoteAddr)

	err := handler(ctx)

	logger.Info(ctx, "request completed", "method", method, "path", path, "remoteaddr", remoteAddr,
		"statuscode", v.StatusCode, "since", time.Since(v.Now).String())

	return err
}
