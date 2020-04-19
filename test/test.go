package test

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "request_count",
			Help: "App Request Count",
		})

	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "my_histogram",
			Help: "This is my histogram",
		})
)

func Count() {
	// rand.Seed(time.Now().Unix())

	prometheus.MustRegister(counter)
	counter.Inc()
	// go func() {
	// 	for {
	// 		counter.Inc()
	// 	}
	// }()
}

func Hist() {
	// rand.Seed(time.Now().Unix())

	histogramVec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "request_latency_seconds",
		Help: "Request latency",
	}, []string{"app_name", "status"},
	)

	prometheus.Register(histogramVec)

	http.Handle("/metrics", newHandlerWithHistogram(promhttp.Handler(), histogramVec))

	prometheus.MustRegister(histogram)
	histogram.Observe(rand.Float64() * 10)
	// go func() {
	// 	for {
	// 		histogram.Observe(rand.Float64() * 10)
	// 	}
	// }()
}

func newHandlerWithHistogram(handler http.Handler, histogram *prometheus.HistogramVec) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		status := http.StatusOK
		//endpoint := req.Host
		serName := "post-srv"

		defer func() {
			histogram.WithLabelValues(fmt.Sprintf("%s %d", serName, status)).Observe(time.Since(start).Seconds())
		}()

		if req.Method == http.MethodGet {
			handler.ServeHTTP(w, req)
			return
		}
		status = http.StatusBadRequest

		w.WriteHeader(status)
	})
}
