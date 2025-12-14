APP_NAME := advent_of_code_2025

.PHONY: all build run test fmt vet clean

all: build

build:
	go build -o bin/$(APP_NAME) .

run:
	go run .

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -rf bin