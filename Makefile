.PHONY: all
all:

.PHONY: lint
lint:
	@go vet `go list ./...`

.PHONY: test
test:
	@go test -v --race --cover -coverprofile=coverage.txt -covermode=atomic ./...
	@go tool cover -html=coverage.txt -o coverage.html
