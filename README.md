# ChatGPT API Usage Exporter

## Introduction

ChatGPT API Usage Exporter is an exporter that monitors ChatGPT API usage.
This exporter retrieves the usage metrics from the ChatGPT API and exports them to Prometheus.

The metrics are daily usage counts (ex. total number of requests) for the ChatGPT API. So the metrics are initialized to 0 at the beginning of the day.

> [!Warning]
>
> This exporter uses the usage API endpoint `https://api.openai.com/v1/usage` that is not officially written in the OpenAI API reference.
> So, the usage API endpoint may be changed or removed in the future.

## Quick Start

To use this exporter, you need to get the OpenAI API key from ChatGPT API dashboard.

1. Set OpenAI API key in the environment variable as `OPENAI_API_KEY`.

   ```bash
   export OPENAI_API_KEY=your_openai_api_key
   ```

2. Run docker-compose.

   ```bash
   docker-compose up
   ```

3. Access `http://localhost:9090` to see the metrics on Prometheus dashboard.

## Metrics

The exporter exports the following metrics:

| Metrics                                  | Type    | Description                                             |
| ---------------------------------------- | ------- | ------------------------------------------------------- |
| chatgpt_api_usage_up                     | gauge   | Was the last query successful                           |
| chatgpt_api_usage_requests_total         | counter | Total number of requests to ChatGPT API per day         |
| chatgpt_api_usage_context_tokens_total   | counter | Total number of context tokens to ChatGPT API per day   |
| chatgpt_api_usage_generated_tokens_total | counter | Total number of generated tokens to ChatGPT API per day |

These metrics except for `chatgpt_api_usage_up ` are labeled by organization ID, organization name, operation and snapshot ID.

> [!Important]
>
> The metrics are initialized to 0 at the beginning of the day.

<details><summary><u>Click to show metrics example</u></summary><p>

```plaintext
# HELP chatgpt_api_usage_context_tokens_total Number of context tokens
# TYPE chatgpt_api_usage_context_tokens_total counter
chatgpt_api_usage_context_tokens_total{operation="completion",organization_id="org-bar",organization_name="foo",snapshot_id="gpt-3.5-turbo-0125"} 47
chatgpt_api_usage_context_tokens_total{operation="completion",organization_id="org-bar",organization_name="foo",snapshot_id="gpt-3.5-turbo-16k-0613"} 203
# HELP chatgpt_api_usage_generated_tokens_total Number of generated tokens
# TYPE chatgpt_api_usage_generated_tokens_total counter
chatgpt_api_usage_generated_tokens_total{operation="completion",organization_id="org-bar",organization_name="foo",snapshot_id="gpt-3.5-turbo-0125"} 20
chatgpt_api_usage_generated_tokens_total{operation="completion",organization_id="org-bar",organization_name="foo",snapshot_id="gpt-3.5-turbo-16k-0613"} 228
# HELP chatgpt_api_usage_requests_total Number of requests
# TYPE chatgpt_api_usage_requests_total counter
chatgpt_api_usage_requests_total{operation="completion",organization_id="org-bar",organization_name="foo",snapshot_id="gpt-3.5-turbo-0125"} 3
chatgpt_api_usage_requests_total{operation="completion",organization_id="org-bar",organization_name="foo",snapshot_id="gpt-3.5-turbo-16k-0613"} 2
# HELP chatgpt_api_usage_up Was the last query successful.
# TYPE chatgpt_api_usage_up counter
chatgpt_api_usage_up 1
```

</p></details>
