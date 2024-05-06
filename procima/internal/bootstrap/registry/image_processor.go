package registry

import (
	"context"
	"github.com/SShlykov/procima/procima/internal/config"
	"github.com/SShlykov/procima/procima/internal/domain/processor"
	loggerPkg "github.com/SShlykov/procima/procima/pkg/logger"
	"github.com/SShlykov/procima/procima/pkg/metrics"
	"sync"
)

func RunImageProcessors(ctx context.Context, logger loggerPkg.Logger, configPath string, routines int, metr metrics.Metrics,
	imageProcessorChan <-chan processor.ImageProcessorItem) error {
	cfg, err := config.LoadProcessorConfig(configPath)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.Run(ctx, logger, cfg.LargestSideLimit, metr, imageProcessorChan)
		}()
	}
	wg.Wait()

	return nil
}
