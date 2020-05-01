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
	req  = http.Request
	resp = req.Response
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
		status = req.Response.Status
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
	status := resp.Status
}

func Output() {
	http.Handle("/metrics", promhttp.Handler())
}
