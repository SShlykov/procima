package app

import (
	"github.com/SShlykov/procima/procima/internal/metrics"
	"github.com/shirou/gopsutil/v3/cpu"
	"runtime"
	"time"
)

const MetricsTimeout = 10 * time.Second

func (app *App) initMetrics() error {
	key := "procima"
	app.metric = metrics.InitMetrics(key)

	return nil
}

func (app *App) RunMetrics() error {
	for {
		select {
		case <-app.ctx.Done():
			return nil
		default:
			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			app.metric.AddMemoryUsage(float64(mem.Alloc))

			percentages, err := cpu.Percent(time.Second, false)
			if err == nil && len(percentages) > 0 {
				app.metric.AddCPUUsage(percentages[0])
			}

			time.Sleep(MetricsTimeout)
		}
	}
}
