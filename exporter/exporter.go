package exporter

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "request_Count",
			Help: "App Request Counter",
		})

	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "my_histogram",
			Help: "This is my histogram",
		})
)

func Count() {
	counterVec := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_count",
		Help: "App Request Count",
	}, []string{"app_name", "method", "endpoint", "http_status"},
	)

	prometheus.Register(counterVec)

	http.Handle("/metrics", newHandlerWithCounter(promhttp.Handler(), counterVec))

	prometheus.MustRegister(counter)
	counter.Inc()
}

func Hist() {
	histogramVec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "request_latency_seconds",
		Help: "Request latency",
	}, []string{"app_name", "method", "endpoint", "http_status"},
	)

	prometheus.Register(histogramVec)

	http.Handle("/metrics", newHandlerWithHistogram(promhttp.Handler(), histogramVec))

	prometheus.MustRegister(histogram)
	histogram.Observe(rand.Float64() * 10)
}

func newHandlerWithCounter(handler http.Handler, counter *prometheus.CounterVec) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		status := req.Response.Status
		endpoint := req.URL.Path
		serName := "post-srv"
		method := req.Method

		defer func() {
			counter.WithLabelValues(serName, method, endpoint, status).Inc()
		}()

		if req.Method == http.MethodGet {
			handler.ServeHTTP(w, req)
			return
		}
	})
}

func newHandlerWithHistogram(handler http.Handler, histogram *prometheus.HistogramVec) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		status := req.Response.Status
		endpoint := req.URL.Path
		serName := "post-srv"
		method := req.Method

		defer func() {
			histogram.WithLabelValues(serName, method, endpoint, status).Observe(time.Since(start).Seconds())
		}()

		if req.Method == http.MethodGet {
			handler.ServeHTTP(w, req)
			return
		}
	})
}
