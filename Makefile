build:
	@go build -o bin/ecommerce cmd/main.go

run: build
	@./bin/ecommerce
