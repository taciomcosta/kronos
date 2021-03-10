KRONOS_VERSION=0.2.0

dev:
	@env ENVIRONMENT=development go run ./cmd/kronosd/main.go
test-unit:
	@go test --tags=unit ./...
test-acceptance:
	@godog test/features
test-all: 
	@echo "======================== UNIT TESTS ======================== "
	@go test --tags=unit ./...
	@echo "======================== ACCEPTANCE TESTS ======================== "
	@godog test/features
codecov:
	@go test -covermode atomic -coverprofile coverage.txt ./...
lint:
	@golangci-lint run
build-any:
	@env ENVIRONMENT=production GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-s -w" -o build/kronosd ./cmd/kronosd
	@env ENVIRONMENT=production GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-s -w" -o build/kronos ./cmd/kronoscli
release-darwin: 
	@env GOOS=darwin GOARCH=amd64 make build-any
	@tar -czvf build/kronos-$(KRONOS_VERSION)-darwin_amd64.tar.gz build/*
release-linux:
	@env GOOS=linux GOARCH=amd64 make build-any
	@cp kronos.service build/kronos.service
	@tar -czvf build/kronos-$(KRONOS_VERSION)-linux_amd64.tar.gz build/*

