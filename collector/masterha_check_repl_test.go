package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func TestUpdateHostStatus(t *testing.T) {
	UpdateHostStatus(prometheus.Labels{"device": "/dev/sda"}, 10)
}
