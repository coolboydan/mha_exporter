package main

import (
	"flag"
	"github.com/coolboydan/mha_exporter/collector"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var addr = flag.String("listen-address", ":9991", "The address to listen on for HTTP request")

func main() {

	flag.Parse()

	collector.UpdateHostStatus()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
