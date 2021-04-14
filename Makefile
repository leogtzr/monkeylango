.DEFAULT_GOAL := build

BIN_FILE=monkey

build:
	@go build -o "${BIN_FILE}"

clean:
	go clean
	rm --force "cp.out"

test:
	go test ./...

check:
	go test ./...

cover:
	go test -coverprofile cp.out ./...
	go tool cover -html=cp.out

run:
	./"${BIN_FILE}"

lint:
	golangci-lint run --enable-all

fmt:
	find . -type f -name '*.go' -exec gofmt -w {} \;
