package collector

import "github.com/prometheus/client_golang/prometheus"

const (
	nameSpace = "chatgpt_api_usage"
)

var commonLabels = []string{"organization_id", "organization_name", "operation", "snapshot_id"}

var (
	up = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "", "up"),
		"Was the last query successful",
		nil,
		nil,
	)
	nRequests = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "", "requests_total"),
		"Total number of requests to ChatGPT API per day",
		commonLabels,
		nil,
	)
	nContextTokensTotal = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "", "context_tokens_total"),
		"Total number of context tokens to ChatGPT API per day",
		commonLabels,
		nil,
	)
	nGeneratedTokensTotal = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "", "generated_tokens_total"),
		"Total number of generated tokens to ChatGPT API per day",
		commonLabels,
		nil,
	)
)

type Metrics struct {
	Up                    int    `json:"up"`
	NRequests             int64  `json:"n_requests"`
	NContextTokensTotal   int64  `json:"n_context_tokens_total"`
	NGeneratedTokensTotal int64  `json:"n_generated_tokens_total"`
	Labels                Labels `json:"labels"`
}

type Labels struct {
	OrganizationID   string `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Operation        string `json:"operation"`
	SnapshotID       string `json:"snapshot_id"`
}

func (l Labels) Values() []string {
	return []string{l.OrganizationID, l.OrganizationName, l.Operation, l.SnapshotID}
}
