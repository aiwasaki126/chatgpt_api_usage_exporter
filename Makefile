.PHONY: setup
setup:
	npm install --save api-spec-converter

.PHONY: convert-oas
convert-oas:
	npx api-spec-converter --from=openapi_3 --to=swagger_2 --syntax=yaml ./openapi/openapi.yml > ./openapi/swagger.yml

.PHONY: generate-response-models
generate-response-models:
	make convert-oas
	docker run --rm -it  --user $(id -u):$(id -g) -v $(HOME):$(HOME) -w $(PWD) quay.io/goswagger/swagger generate model -f ./openapi/swagger.yml -t client -m response

.PHONY: build
build:
	CGO_ENABLED=0 go build -o main

.PHONY: docker-build
docker-build:
	docker build --no-cache -t chatgpt-api-usage-exporter .
