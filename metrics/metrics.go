package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Count = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_count",     //Unique id, can't repeat Register(), can Unregister()
		Help: "App Request Count", //Description of this Counter
	},
		[]string{"app_name", "method", "endpoint", "http_status"},
	)
	Latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "request_latency_seconds",
		Help:    "Request latency",
		Buckets: prometheus.LinearBuckets(0, 1, 10), //There are 20 first barrels, 5 intervals for each barrel, 5 barrels in total. So 20, 25, 30, 35, 40
	},
		[]string{"app_name", "endpoint"},
	)
)

func init() {
	//Cannot register Metrics with the same Name more than once
	//MustRegister registration failure will directly panic(), if you want to capture error, it is recommended to use Register()
	prometheus.MustRegister(Count)
	prometheus.MustRegister(Latency)
}

func Timer() {
	timer := prometheus.NewTimer(Latency)
	defer timer.ObserveDuration()
}

func PostCount() {
	Count.With(prometheus.Labels{"app_name": "post-srv", "method": http.Handle, "endpoint": http.Request, "http_status": http.ResponseWriter}).Inc()
}

func PostHist() {
	Latency.With(prometheus.Labels{"app_name": "post-srv", "endpoint": http.Request}).Inc()
}

func Output() {
	http.Handle("/metrics", promhttp.Handler())
}
