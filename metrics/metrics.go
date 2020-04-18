package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	MyTestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		//Because the Name cannot be duplicate, the recommended rule is: "department Name business Name module Name scalar Name type"
		Name: "request_count",     //Unique id, can't repeat Register(), can Unregister()
		Help: "App Request Count", //Description of this Counter
	})
	MyTestHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "request_latency_seconds",
		Help:    "Request latency",
		Buckets: prometheus.LinearBuckets(20, 5, 5), //There are 20 first barrels, 5 intervals for each barrel, 5 barrels in total. So 20, 25, 30, 35, 40
	})
	MyTestSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "my_test_summary",
		Help:       "my test summary",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}, //Return five, nine, nine
	})
)

func Collect() {

	//Cannot register Metrics with the same Name more than once
	//MustRegister registration failure will directly panic(), if you want to capture error, it is recommended to use Register()
	prometheus.MustRegister(MyTestCounter)
	prometheus.MustRegister(MyTestHistogram)
	prometheus.MustRegister(MyTestSummary)

	// go func() {
	// 	var i float64
	// 	for {
	// 		i++
	// 		MyTestCounter.Add(10000)                                                  //Constant added each time
	// 		MyTestHistogram.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10) //Observe a quantity of 18 - 42 at a time
	// 		MyTestSummary.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)

	// 		time.Sleep(time.Second)
	// 	}
	// }()
	// http.Handle("/metrics", promhttp.Handler())
	// log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))//Multiple processes cannot listen to the same port
}

func Timer() {
	timer := prometheus.NewTimer(MyTestHistogram)
	defer timer.ObserveDuration()
}

func Output() {
	http.Handle("/metrics", promhttp.Handler())
}

// import (
// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// 	"time"
// )

// var (
// 	counter = prometheus.NewCounter(
// 		prometheus.CounterOpts{
// 			Name: "request_count",
// 		})

// 	latency = prometheus.Histogram(
// 		prometheus.HistogramOpts{
// 			Name:    "request_latency_seconds",
// 			Buckets: prometheus.LinearBuckets(0, 10, 20),
// 		})
// )

// func init() {
// 	prometheus.MustRegister(counter)
// 	prometheus.MustRegister(latency)

// }
// func timer() {
// 	timer := prometheus.NewTimer(myHistogram)
// 	defer timer.ObserveDuration()
// 	// Do actual work.
// }

// func metrics() {

// }
