package mid

import (
	"context"
	"fmt"
	"github.com/matheusgcoppi/service/app/api/metrics"
	"runtime/debug"
)

// Panics recovers from panics and converts the panic to an error so it is
// reported in Metrics and handled in Errors.
func Panics(ctx context.Context, handler Handler) (err error) {

	// Defer a function to recover from a panic and set the err return
	// variable after the fact.
	// Defer func will run after the return
	defer func() {
		if rec := recover(); rec != nil {
			trace := debug.Stack()
			err = fmt.Errorf("PANIC [%v] TRACE[%s]", rec, string(trace))

			metrics.AddPanics(ctx)
		}
	}()

	return handler(ctx)
}
