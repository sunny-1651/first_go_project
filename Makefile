build:
	@go build -o bin/fs

run: build
	@./bin/fs

test:
	@go test ./... -v


# to add module and requirements
tidy:
	@go mod tidy
