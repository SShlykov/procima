package app

import (
	"github.com/SShlykov/procima/procima/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
	"runtime"
	"time"
)

const MetricsTimeout = 10 * time.Second

func (app *App) initMetrics() error {
	key := "procima"

	counter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: key,
			Name:      "gin_request_counter",
			Help:      "When request is made, this counter is incremented by 1",
		},
	)

	durRecorder := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: key,
			Name:      "gin_response_time",
			Help:      "A histogram of the response time in microseconds.",
		},
	)

	parseRecorder := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: key,
			Name:      "image_perser_time",
			Help:      "A histogram of the response time in microseconds.",
		},
	)

	memoryUsage := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: key,
			Name:      "system_memory_usage",
			Help:      "System memory usage",
		},
	)

	cpuUsage := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: key,
			Name:      "system_cpu_usage",
			Help:      "System CPU usage",
		},
	)

	app.metric = metrics.NewMetrics(counter, durRecorder, parseRecorder, memoryUsage, cpuUsage)

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
