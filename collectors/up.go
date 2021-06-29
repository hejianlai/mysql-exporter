package collectors

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type UPcollector struct {
	*baseCollector
	desc *prometheus.Desc
}

func NewUpCollector(db *sql.DB) *UPcollector {
	desc := prometheus.NewDesc("mysql_up", "mysql health", nil, nil)
	return &UPcollector{newBaseCollector(db), desc}
}
func (c *UPcollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}
func (c *UPcollector) Collect(metrics chan<- prometheus.Metric) {
	up := 1
	if err := c.db.Ping(); err != nil {
		up = 0
	}
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.GaugeValue, float64(up))
}
