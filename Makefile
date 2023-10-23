build:
	@go build -o bin/al

run: build
	@./bin/al

test:
	@go test ./... -v