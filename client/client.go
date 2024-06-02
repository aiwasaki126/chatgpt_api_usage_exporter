package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/aiwasaki126/chat_gpt_api_usage_exporter/collector"
)

type client struct {
	baseURL url.URL
	apiKey  string
}

func New(targetURL, apiKey string) (*client, error) {
	u, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}
	return &client{
		baseURL: *u,
		apiKey:  apiKey,
	}, nil
}

func (c *client) GetUsageMetrics(ctx context.Context) ([]*collector.Metrics, error) {
	client := http.Client{
		Transport: http.DefaultTransport,
		Timeout:   time.Second * 10,
	}
	req, err := c.setupRequest(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Failed to get usage metrics from API
		return nil, handleErrorResponse(resp)
	}
	converter := newMetricConverter()
	return converter.extractUsageMetrics(resp)
}

func (c *client) setupRequest(ctx context.Context) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL.String(), nil)
	if err != nil {
		return nil, err
	}
	// Set the Authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	// Add the current date to the query
	query := req.URL.Query()
	query.Add("date", time.Now().Format(time.DateOnly))
	if _, err := url.ParseQuery(query.Encode()); err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()

	return req, nil
}
