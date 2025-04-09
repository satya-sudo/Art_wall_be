APP_NAME := gallery-server

.PHONY: all build run clean tidy

all: build

## 🔨 Build the binary
build:
	go build -o bin/$(APP_NAME) main.go

## ▶️ Run the server (uses .env if needed)
run:
	go run main.go

## 🧹 Clean build files
clean:
	rm -rf bin/

## 🧽 Tidy up dependencies
tidy:
	go mod tidy
