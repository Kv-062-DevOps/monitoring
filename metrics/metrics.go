package metrics

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
	req = http.Request
)

func CountRegister() {
	prometheus.Register(CounterVec)
}

func HistRegister() {
	prometheus.Register(HistogramVec)
}

func Init() {
	start := time.Now()
	status := ""
	endpoint := req.URL.Path
	serName := "post-srv"
	method := req.Method

}

func CountCollect() {
	Init()
	defer func() {
		CounterVec.WithLabelValues(serName, method, endpoint, status).Inc()
	}()
}

func HistCollect() {
	Init()
	defer func() {
		HistogramVec.WithLabelValues(serName, endpoint).Observe(time.Since(start).Seconds())
	}()

}

func StatusCollect() {
	status := Status
}

func Output() {
	http.Handle("/metrics", promhttp.Handler())
}

// func newHandlerWithHistogram(handler http.Handler, histogram *prometheus.HistogramVec) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 		//start := time.Now()
// 		// status := req.Response.Status
// 		// endpoint := req.URL.Path
// 		// serName := "post-srv"
// 		// method := req.Method

// 		defer func() {
// 			histogram.WithLabelValues(serName, method, endpoint, status).Observe(time.Since(start).Seconds())
// 		}()

// if req.Method == http.MethodGet {
// 	handler.ServeHTTP(w, req)
// 	return
// }
// status = http.StatusBadRequest

// w.WriteHeader(status)
// 	})
// }
