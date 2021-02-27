package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
)

type httpMetrics struct {
	counterMetrics map[string]prometheus.Counter
}

var HttpMetrics *httpMetrics
var Registry *prometheus.Registry

func init() {
	HttpMetrics = &httpMetrics{
		counterMetrics: make(map[string]prometheus.Counter),
	}

	Registry = prometheus.NewRegistry()
}

func checkIfCounterMetricExist(counter prometheus.Counter) bool {
	_, ok := HttpMetrics.counterMetrics[counter.Desc().String()]
	return ok
}

func (metrics httpMetrics) AddCounterMetric(name string, help string) {

	var counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
	)

	if checkIfCounterMetricExist(counter) == false { //TODO refactor if different requests will be huge
		HttpMetrics.counterMetrics[counter.Desc().String()] = counter
		Registry.MustRegister(HttpMetrics.counterMetrics[counter.Desc().String()])
	}

	HttpMetrics.counterMetrics[counter.Desc().String()].Inc()
}

func (metrics httpMetrics) RemoveCounterMetric(counter prometheus.Counter) {
	Registry.Unregister(counter)
	delete(HttpMetrics.counterMetrics, counter.Desc().String())
}
