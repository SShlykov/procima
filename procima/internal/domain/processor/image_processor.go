package processor

import (
	"context"
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
)

func ImageProcessor(ctx context.Context, logger loggerPkg.Logger) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
