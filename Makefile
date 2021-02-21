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
codecov:
	@go test -covermode atomic -coverprofile coverage.txt ./...
lint:
	@golangci-lint run
build-any:
	@env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-s -w" -o build/kronosd ./cmd/kronosd
	@env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-s -w" -o build/kronos ./cmd/kronoscli
release-darwin: 
	@cp scripts/install.bash build/install.bash
	@cp scripts/com.taciomcosta.kronos.plist build/com.taciomcosta.kronos.plist
	@env GOOS=darwin GOARCH=amd64 make build-any
	@tar -czvf build/kronos-$(KRONOS_VERSION)-darwin_amd64.tar.gz build/*
