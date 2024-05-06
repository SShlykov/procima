package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	Request()
	ResponseDuration(duration float64)
	ImageParseDuration(duration float64)
	AddMemoryUsage(value float64)
	AddCPUUsage(value float64)
	Register(reg *prometheus.Registry)
}

type metrics struct {
	counter            prometheus.Counter
	responseDuration   prometheus.Histogram
	imageParseDuration prometheus.Histogram
	memoryUsage        prometheus.Counter
	cpuUsage           prometheus.Counter
}

func NewMetrics(counter prometheus.Counter, responseDuration prometheus.Histogram,
	imageParseDuration prometheus.Histogram, memoryUsage prometheus.Counter, cpuUsage prometheus.Counter) Metrics {
	return &metrics{counter: counter, responseDuration: responseDuration, imageParseDuration: imageParseDuration,
		memoryUsage: memoryUsage, cpuUsage: cpuUsage}
}

func (m *metrics) Register(reg *prometheus.Registry) {
	reg.MustRegister(m.counter)
	reg.MustRegister(m.responseDuration)
	reg.MustRegister(m.imageParseDuration)
	reg.MustRegister(m.memoryUsage)
	reg.MustRegister(m.cpuUsage)
}

func (m *metrics) AddMemoryUsage(value float64) {
	m.memoryUsage.Add(value)
}

func (m *metrics) AddCPUUsage(value float64) {
	m.cpuUsage.Add(value)
}

func (m *metrics) ImageParseDuration(duration float64) {
	m.imageParseDuration.Observe(duration)
}

func (m *metrics) ResponseDuration(duration float64) {
	m.responseDuration.Observe(duration)
}

func (m *metrics) Request() {
	m.counter.Add(1)
}
