package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_req_count_ping",
		Help: "No of request handled by Ping handler",
	},
)

var gauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "gauge_example",
	Help: "Gauge example",
})

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	gauge.Inc()

	fmt.Fprintf(w, "pong")
}

func gnip(w http.ResponseWriter, req *http.Request) {
	gauge.Dec()
	fmt.Fprintf(w, "gnop")
}

func main() {
	prometheus.MustRegister(pingCounter)
	prometheus.MustRegister(gauge)

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/reverse-ping", gnip)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("starting server...")
	http.ListenAndServe(":2112", nil)
}
