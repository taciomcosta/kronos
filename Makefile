KRONOS_VERSION=0.1.0

all:
	@go run ./cmd/kronosd/main.go
test-unit:
	@go test --tags=unit ./...
test-acceptance:
	@godog test/features
test-all: 
	@echo "======================== UNIT TESTS ======================== "
	@go test --tags=unit ./...
	@echo "======================== ACCEPTANCE TESTS ======================== "
	@godog test/features
cover:
	@go test ./... -v -cover
lint:
	@golangci-lint run
build-darwin:
	@env GOOS=darwin GOARCH=amd64 make build-any
build-any:
	@env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-s -w" -o build/kronosd ./cmd/kronosd
	@env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-s -w" -o build/kronos ./cmd/kronoscli
	@tar -czvf build/kronos-$(KRONOS_VERSION)-$(GOOS)_$(GOARCH).tar.gz build/*

