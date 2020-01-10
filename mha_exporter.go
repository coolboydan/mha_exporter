package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var addr = flag.String("listen-address", ":9991", "The address to listen on for HTTP request")

var (
	uniformDomain = flag.Float64("uniform.domain", 0.0002, "The domain for the uniform distribution.")
	normDomain    = flag.Float64("normal.domain", 0.0002, "The domain for the normal distribution.")
	normMean      = flag.Float64("normal.mean", 0.00001, "The mean for the normal distribution.")

	opsQueued = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "ops_queued",
		Help:      "Number of bolob storage operations waiting to process.",
	}, []string{"name"})

	taskCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Subsystem: "work_pool",
		Name:      "completed_task_total",
		Help:      "Total number of tasks completed.",
	})

	temps = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "pond_temperature_celsius",
		Help:       "The temperature of the frog pond.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})

	rpcDurationsHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "rpc_durations_histogram_cesconds",
		Help:    "RPC late distr",
		Buckets: []float64{1, 2, 5, 10, 20, 60},
	})
)

func init() {
	prometheus.MustRegister(opsQueued)
	prometheus.MustRegister(taskCounter)
	prometheus.MustRegister(temps)
	prometheus.MustRegister(rpcDurationsHistogram)
}

func main() {

	flag.Parse()
	go func() {
		for {

			opsQueued.With(prometheus.Labels{"device": "/dev/sda"}).Set(time.Millisecond.Seconds())

			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			taskCounter.Inc()

			time.Sleep(time.Second * 4)
		}
	}()

	go func() {
		for {
			temps.Observe(rand.Float64() * *uniformDomain)

			time.Sleep(time.Second * 4)
		}
	}()

	go func() {
		for {

			temps.Observe(rand.Float64() * *uniformDomain)
			time.Sleep(time.Second * 4)
		}
	}()

	go func() {
		for {
			rpcDurationsHistogram.Observe((rand.NormFloat64() * *normDomain) + *normMean)

			time.Sleep(time.Second * 4)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
