package client

import (
	"encoding/json"
	"net/http"

	"github.com/aiwasaki126/chat_gpt_api_usage_exporter/client/response"
	"github.com/aiwasaki126/chat_gpt_api_usage_exporter/collector"
)

type metricsConverter struct {
	body *response.UsageResponse
}

func newMetricConverter() *metricsConverter {
	return &metricsConverter{}
}

func (m *metricsConverter) extractUsageMetrics(resp *http.Response) ([]*collector.Metrics, error) {
	if err := m.validateResponse(resp); err != nil {
		return nil, err
	}
	if m.body == nil {
		return nil, nil
	}
	rawMetrics := extractRawMetrics(m.body.Data)
	return m.convertDataToMetrics(rawMetrics)
}

func (m *metricsConverter) validateResponse(resp *http.Response) error {
	var body response.UsageResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return err
	}
	if err := body.Validate(nil); err != nil {
		return err
	}
	m.body = &body
	return nil
}

func (m *metricsConverter) convertDataToMetrics(data []*rawMetrics) ([]*collector.Metrics, error) {
	metricsMap := make(map[collector.Labels]collector.Metrics, 0)
	for _, d := range data {
		label := collector.Labels{
			OrganizationID:   d.OrganizationID,
			OrganizationName: d.OrganizationName,
			Operation:        d.Operation,
			SnapshotID:       d.SnapshotID,
		}
		nRequests := int64(0)
		nContextTokensTotal := int64(0)
		nGeneratedTokensTotal := int64(0)

		metrics, ok := metricsMap[label]
		if ok {
			nRequests = metrics.NRequests
			nContextTokensTotal = metrics.NContextTokensTotal
			nGeneratedTokensTotal = metrics.NGeneratedTokensTotal
		}
		metricsMap[label] = collector.Metrics{
			Up:                    1,
			NRequests:             nRequests + d.NRequests,
			NContextTokensTotal:   nContextTokensTotal + d.NContextTokensTotal,
			NGeneratedTokensTotal: nGeneratedTokensTotal + d.NGeneratedTokensTotal,
			Labels:                label,
		}
	}
	metrics := make([]*collector.Metrics, 0, len(metricsMap))
	for _, m := range metricsMap {
		metrics = append(metrics, &m)
	}
	return metrics, nil
}

type rawMetrics struct {
	NRequests             int64 `json:"n_requests"`
	NContextTokensTotal   int64 `json:"n_context_tokens"`
	NGeneratedTokensTotal int64 `json:"n_generated_tokens"`
	collector.Labels
}

func extractRawMetrics(data []*response.Data) []*rawMetrics {
	metricsMaterials := make([]*rawMetrics, 0, len(data))
	for _, d := range data {
		metricsMaterials = append(metricsMaterials, &rawMetrics{
			NRequests:             *d.NRequests,
			NContextTokensTotal:   *d.NContextTokensTotal,
			NGeneratedTokensTotal: *d.NGeneratedTokensTotal,
			Labels: collector.Labels{
				OrganizationID:   *d.OrganizationID,
				OrganizationName: *d.OrganizationName,
				Operation:        *d.Operation,
				SnapshotID:       *d.SnapshotID,
			},
		})
	}
	return metricsMaterials

}
