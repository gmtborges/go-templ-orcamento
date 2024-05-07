# Makefile for building and deploying a Go application to a Debian amd64 server.

GOARCH := amd64
GOOS := linux

.PHONY: all build css-minify templ templ-proxy

all: css-minify build

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/app .

css-minify:
	npx tailwindcss -i ./assets/input.css -o ./public/output.css --minify

css:
	npx tailwindcss --watch -i ./assets/input.css -o ./public/output.css

templ:
	templ generate

templ-proxy:
	templ generate --watch --proxy=http://127.0.0.1:3000

clean:
	rm -f bin/$(APP_NAME)

