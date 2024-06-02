package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aiwasaki126/chat_gpt_api_usage_exporter/client"
	"github.com/aiwasaki126/chat_gpt_api_usage_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	targetURL = "https://api.openai.com/v1/usage"
)

var port = flag.Int("port", 8080, "port number to listen")

func main() {
	// setup collector
	mustGetEnv := func(key string) string {
		v := os.Getenv(key)
		if v == "" {
			log.Fatalf("missing %s", key)
		}
		return v
	}
	apiClient, err := client.New(targetURL, mustGetEnv("OPENAI_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	// register collector
	c := collector.NewCollector(apiClient)
	prometheus.MustRegister(c)

	// start exporter
	flag.Parse()
	log.Printf("listening on :%v", *port)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), nil))
}
