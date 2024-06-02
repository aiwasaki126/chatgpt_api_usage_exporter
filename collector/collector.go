package collector

import (
	"context"
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	retriever Retriever
}

func NewCollector(retriever Retriever) *Collector {
	return &Collector{
		retriever: retriever,
	}
}

func (c Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	ch <- nRequests
	ch <- nContextTokensTotal
	ch <- nGeneratedTokensTotal
}

func (c Collector) Collect(ch chan<- prometheus.Metric) {
	metrics, err := c.retriever.GetUsageMetrics(context.Background())
	if err != nil {
		log.Printf("failed to get usage metrics: %v", err)
		ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 0)
		return
	}
	// Was the last query successful
	ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 1)
	for _, metric := range metrics {
		// Main metrics
		ch <- prometheus.MustNewConstMetric(
			nRequests,
			prometheus.CounterValue,
			float64(metric.NRequests),
			metric.Labels.Values()...)
		ch <- prometheus.MustNewConstMetric(
			nContextTokensTotal,
			prometheus.CounterValue,
			float64(metric.NContextTokensTotal),
			metric.Labels.Values()...)
		ch <- prometheus.MustNewConstMetric(
			nGeneratedTokensTotal,
			prometheus.CounterValue,
			float64(metric.NGeneratedTokensTotal),
			metric.Labels.Values()...)
	}
}
