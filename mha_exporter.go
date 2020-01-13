package main

import (
	"flag"
	"github.com/coolboydan/mha_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var addr = flag.String("listen-address", ":9991", "The address to listen on for HTTP request")

func main() {

	flag.Parse()

	collector.UpdateHostStatus(prometheus.Labels{"name": "app1", "host": "192.186.1.1:3306", "status": "health"}, 0)
	collector.UpdateHostStatus(prometheus.Labels{"name": "app1", "host": "192.186.1.1:3306", "status": "role"}, 2)
	collector.UpdateHostStatus(prometheus.Labels{"name": "app1", "host": "192.186.1.2:3306", "status": "role"}, 3)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
