package collector

import "github.com/prometheus/client_golang/prometheus"

var opsQueued = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "",
	Subsystem: "blob_storage",
	Name:      "ops_queued",
	Help:      "Number of bolob storage operations waiting to process.",
}, []string{"name"})

func UpdateHostStatus(labels prometheus.Labels, status float64) {
	opsQueued.With(labels).Set(status)
}

func init() {
	prometheus.MustRegister(opsQueued)
}
