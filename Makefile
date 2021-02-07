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
	@golint ./...
