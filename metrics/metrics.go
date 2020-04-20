package metrics

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
			Name: "app_name",
			Help: "request_app_counter",
		})

	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "golang",
			Name:      "my_histogram",
			Help:      "This is my histogram",
		})
)

func Count() {
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

func newHandlerWithHistogram(handler http.Handler, histogram *prometheus.HistogramVec) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		status := req.Response.Status
		endpoint := req.URL.Path
		serName := "post-srv"
		method := req.Method

		defer func() {
			histogram.WithLabelValues(serName, method, endpoint, status).Observe(time.Since(start).Seconds())
		}())

	// 	if req.Method == http.MethodGet {
	// 		handler.ServeHTTP(w, req)
	// 		return
	// 	}
	// 	status = http.StatusBadRequest

	// 	w.WriteHeader(status)
	// })
}
