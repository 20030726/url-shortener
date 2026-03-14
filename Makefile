.PHONY: all frontend build run test clean

all: build

## Install frontend dependencies and compile TypeScript
frontend:
	cd frontend && npm install && npm run build

## Build the Go binary (frontend must be built first)
build: frontend
	go build -o url-shortener .

## Run the server (builds everything first)
run: build
	./url-shortener

## Run Go tests
test:
	go test ./...

## Remove build artifacts
clean:
	rm -f url-shortener
	rm -rf frontend/dist
