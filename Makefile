# Makefile for building and deploying a Go application to a Debian amd64 server.

GOARCH := amd64
GOOS := linux

.PHONY: all build css-minify

all: css-minify build

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/app .

css-minify:
	tailwindcss -i ./assets/input.css -o ./public/output.css --minify

css:
	tailwindcss --watch -i ./assets/input.css -o ./public/output.css

clean:
	rm -f bin/$(APP_NAME)

