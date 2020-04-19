package metrics

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Count = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_count",
		Help: "App Request Count",
	},
		[]string{"app_name", "method", "endpoint", "http_status"},
	)
	Latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "request_latency_seconds",
		Help: "Request latency",
	},
		[]string{"app_name", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(Count)
	prometheus.MustRegister(Latency)
}

func MeasureTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer r.Body.Close()
		code := http.StatusInternalServerError

		defer func() { // Make sure we record a status.
			duration := time.Since(start)
			Latency.WithLabelValues(fmt.Sprintf("%d", code)).Observe(duration.Seconds())
		}()
	}
}

func Collect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//start := time.Now()
		defer r.Body.Close()
		//code := http.StatusInternalServerError

		defer func() { // Make sure we record a status.
			//duration := time.Since(start)
			Count.With(prometheus.Labels{"app_name": "post-srv", "method": r.Method,
				"endpoint": r.Host, "http_status": r.Response.Status}).Inc()
		}()
	}
}

// func StartTimer() {
// 	http.Request.StartTime = time.Now()
// }

// func StartTime() {
// 	var Start = time.Now()
// }

// func MeasureTime() {
// 	var Start = time.Now()
// 	Latency.Observer(time.Since(Start).Seconds())
// }

// func PostCount() {
// 	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
// 		Count.With(prometheus.Labels{"app_name": "post-srv", "method": r.Method,
// 			"endpoint": r.Host, "http_status": r.Response.Status}).Inc()
// 	})
// }

// func PostHist() {
// 	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
// 		Latency.With(prometheus.Labels{"app_name": "post-srv", "endpoint": r.Host})
// 	})
// }

func Output() {
	http.Handle("/metrics", promhttp.Handler())
}
