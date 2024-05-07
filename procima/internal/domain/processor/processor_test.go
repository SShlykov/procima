package processor

import (
	"context"
	loggerPkg "github.com/SShlykov/procima/procima/pkg/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
	"image"
	"testing"
	"time"
)

type MetricsMock struct{}

func (m *MetricsMock) ImageParseDuration(_ float64) {}

func TestRun_MultipleGoroutines(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger := loggerPkg.SetupLogger("debug", "pretty", "test", "no-host")
	metric := &MetricsMock{}
	procChan := make(chan ImageProcessorItem, 10)

	go Run(ctx, logger, 100, metric, procChan)

	for range 10 {
		resultChan := make(chan ImageResult, 1)
		procChan <- ImageProcessorItem{
			Channel:   resultChan,
			Operation: "rotate:90deg",
			Img:       image.NewRGBA(image.Rect(0, 0, 100, 100)),
		}

		go func() {
			select {
			case res := <-resultChan:
				assert.NoError(t, res.Err)
				assert.NotNil(t, res.Res)
			case <-ctx.Done():
				t.Errorf("Context canceled before receiving the result")
			}
		}()
	}

	time.Sleep(1 * time.Second) // Дать время на обработку
}
