package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	CounterVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_count",
		Help: "App Request Count",
	}, []string{"app_name", "method", "endpoint", "http_status"},
	)
	HistogramVec = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "request_latency_seconds",
		Help: "Request latency",
	}, []string{"app_name", "endpoint"},
	)
)

func Count() {
	prometheus.Register(CounterVec)
}

func HistRegister() {
	prometheus.Register(HistogramVec)
}

func Output() {
	http.Handle("/metrics", promhttp.Handler())
}
