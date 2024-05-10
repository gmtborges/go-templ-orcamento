# Makefile for building and deploying a Go application to a Debian amd64 server.

GOARCH := amd64
GOOS := linux

.PHONY: all build css-minify templ templ-proxy migrate

all: css-minify templ build

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/app .

css-minify:
	npx tailwindcss -i ./assets/input.css -o ./static/output.css --minify

css:
	npx tailwindcss --watch -i ./assets/input.css -o ./static/output.css

templ:
	templ generate

templ-proxy:
	templ generate --watch --proxy=http://127.0.0.1:3000

migrate:
	go run ./migrations
	
clean:
	rm -f bin/$(APP_NAME)

