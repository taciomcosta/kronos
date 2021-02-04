all:
	@go run ./cmd/kronosd/main.go
test:
	@go test ./...
lint:
	golint ./...
cover:
	@go test ./... -cover
