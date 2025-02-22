build:
	@go build -o bin/udemygo cmd/web/main.go
run: build
	@bin/udemygo
