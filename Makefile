GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOGET=$(GOCMD) get
PROJECT_FILES=$(go list ./... | grep -v /vendor/)
BINARY_NAME=ess-atomic-swap

## Develop
run_development: build_development start_development

build_development:
	$(GOGET) github.com/codegangsta/gin
	$(GOGET) -u github.com/golangci/golangci-lint/cmd/golangci-lint

start_development:
	gin --appPort=${ESS_ATOMIC_SWAP_APP_PORT} --port=${ESS_ATOMIC_SWAP_DEV_PORT} run main.go

## Production
run_production:	build_production start_production

build_production:
	$(GOBUILD) .

start_production:
	./$(BINARY_NAME)

## Linters
run_golangci_lint:
	golangci-lint run --config ./.golangci.yml ./...

## Tests
run_unit_tests:
	go test $(PROJECT_FILES)
