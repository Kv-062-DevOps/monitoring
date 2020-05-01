package exporter

import (
	"net/http"
	"time"

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
	status, endpoint, serName, method string = "", "", "", ""
)

func RegisterMetrics() {
	prometheus.Register(CounterVec)
	prometheus.Register(HistogramVec)
}

func InitCounter() {
	http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start = time.Now()
		serName = "post-srv"
		method = req.Method
		endpoint = req.URL.Path
		status = ""
	})
}

func InitHist() {
	http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		serName = "post-srv"
		endpoint = req.URL.Path
	})
}

func CollectCount() {
	InitCounter()
	CounterVec.WithLabelValues(serName, method, endpoint, status).Inc()
}

func CollectHist() {
	InitHist()
	HistogramVec.WithLabelValues(serName, endpoint).Observe(time.Since(start).Seconds())
}

func StatusCollect() {
	http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		status = req.Response.Status
	})

}

func Output() {
	http.Handle("/metrics", promhttp.Handler())
}
