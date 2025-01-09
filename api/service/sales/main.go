package sales

import (
	"context"
	"os"
)

func main() {
	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "msg", err)
		os.Exit(1)
	}

}

func run(ctx context.Context log *logger.Logger) error {

}
