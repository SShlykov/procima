package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

func InitMetrics(key string) Metrics {
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

	return NewMetrics(counter, durRecorder, parseRecorder, memoryUsage, cpuUsage)
}
