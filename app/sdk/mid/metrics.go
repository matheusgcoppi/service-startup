package mid

import (
	"context"
	"github.com/matheusgcoppi/service/app/sdk/metrics"
)

func Metrics(ctx context.Context, handler Handler) error {
	ctx = metrics.Set(ctx)

	err := handler(ctx)

	n := metrics.AddRequests(ctx)
	if n%1000 == 0 {
		metrics.AddGoroutines(ctx)
	}

	if err != nil {
		metrics.AddErrors(ctx)
	}

	return err
}
