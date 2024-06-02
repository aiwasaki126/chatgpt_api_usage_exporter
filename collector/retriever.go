package collector

import "context"

type Retriever interface {
	GetUsageMetrics(context.Context) ([]*Metrics, error)
}
