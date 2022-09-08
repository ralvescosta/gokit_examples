package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

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

func sysMetrics() {
	//mem
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Println("go_memstats_sys_bytes Number of bytes obtained from system.", float64(m.Sys))
	log.Println("go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.", float64(m.TotalAlloc))
	log.Println("go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.", float64(m.HeapAlloc))
	log.Println("go_memstats_frees_total Total number of frees.", float64(m.Frees))
	log.Println("go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.", float64(m.GCSys))
	log.Println("go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.", float64(m.HeapIdle))
	log.Println("go_memstats_heap_inuse_bytes Number of heap bytes that are in use.", float64(m.HeapInuse))
	log.Println("go_memstats_heap_objects Number of allocated objects.", float64(m.HeapObjects))
	log.Println("go_memstats_heap_released_bytes Number of heap bytes released to OS.", float64(m.HeapReleased))
	log.Println("go_memstats_heap_sys_bytes Number of heap bytes obtained from system.", float64(m.HeapSys))
	log.Println("go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.", float64(m.LastGC))
	log.Println("go_memstats_lookups_total Total number of pointer lookups.", float64(m.Lookups))
	log.Println("go_memstats_mallocs_total Total number of mallocs.", float64(m.Mallocs))
	log.Println("go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.", float64(m.MCacheInuse))
	log.Println("go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.", float64(m.MCacheSys))
	log.Println("go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.", float64(m.MSpanInuse))
	log.Println("go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.", float64(m.MSpanSys))
	log.Println(" go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.", float64(m.NextGC))
	log.Println("go_memstats_other_sys_bytes Number of bytes used for other system allocations.", float64(m.OtherSys))
	log.Println("go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.", float64(m.StackSys))
	log.Println("go_memstats_gc_completed_cycle Number of GC cycle completed.", float64(m.NumGC))
	log.Println("go_memstats_gc_pause_total Number of GC-stop-the-world caused in Nanosecond", float64(m.PauseTotalNs))

	//os
	log.Println("go_threads Number of OS threads created.", float64(runtime.NumCPU()))
	log.Println("go_cgo Number of CGO.", float64(runtime.NumCgoCall()))
	log.Println("go_goroutines Number of goroutines.", float64(runtime.NumGoroutine()))
}

func main() {
	prometheus.MustRegister(pingCounter)
	prometheus.MustRegister(gauge)

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/reverse-ping", gnip)
	http.Handle("/metrics", promhttp.Handler())
	go sysMetrics()

	log.Println("starting server...")
	http.ListenAndServe(":2112", nil)
}
