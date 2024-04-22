# Makefile for building and deploying a Go application to a Debian amd64 server.

APP_NAME := myapp
SERVER := gcloud-vm
DEST_PATH := /home/gcloud
GOARCH := amd64
GOOS := linux

.PHONY: all build deploy

all: build deploy

build:
	@echo "Building the Go application..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$(APP_NAME) ./cmd

deploy:
	@echo "Deploying the application to the server..."
	scp $(APP_NAME) $(SERVER):$(DEST_PATH)

clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

