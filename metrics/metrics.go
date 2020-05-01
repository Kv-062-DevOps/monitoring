package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// CounterVec vector, to count amount of requests
	CounterVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_count",
		Help: "App Request Count",
	}, []string{"app_name", "method", "endpoint", "http_status"},
	)
	// HistogramVec vector, to measure time
	HistogramVec = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "request_latency_seconds",
		Help: "Request latency",
	}, []string{"app_name", "endpoint"},
	)
)

// RegMetrics func for registering metrics, so they can be exposed
func RegMetrics() {
	prometheus.Register(CounterVec)
	prometheus.Register(HistogramVec)
}

// Output func which provides handler for exposing metrics via an HTTP server,
// "/metrics" is the usual endpoint for that
func Output() {
	http.Handle("/metrics", promhttp.Handler())
}
