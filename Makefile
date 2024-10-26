GOBIN ?= $$(go env GOPATH)/bin

.PHONY: check-coverage lint

dev:
	wails dev

mocks:
	mockery && \
	sed -i 's v2/internal/frontend v2/pkg/runtime ' 'internal/interfaces/mocks/mock_WailsRuntime.go'

test:
	go test ./...

.PHONY: install-go-test-coverage
install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

.PHONY: check-coverage
check-coverage: install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.testcoverage.yaml

lint:
	golangci-lint run