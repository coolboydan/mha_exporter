package collector

import "github.com/prometheus/client_golang/prometheus"

var hostVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "host",
	Subsystem: "status",
	Name:      "name",
	Help:      "host status.",
}, []string{"name", "host", "status"})

func UpdateHostStatus(labels prometheus.Labels, status float64) {
	hostVec.With(labels).Set(status)
}

func init() {
	prometheus.MustRegister(hostVec)
}
