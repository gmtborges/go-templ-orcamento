# Makefile for building and deploying a Go application to a Debian amd64 server.

GOARCH := amd64
GOOS := linux

build: css-minify templ
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/app .

css-minify:
	npx tailwindcss -i ./assets/input.css -o ./static/output.css --minify

templ:
	templ generate

css:
	@npx tailwindcss --watch -i ./assets/input.css -o ./static/output.css

proxy:
	templ generate --watch --proxy=http://127.0.0.1:3000

migrate:
	@go run ./cmd/migrate up

seed:
	@go run ./cmd/seed
	
clean:
	rm -f bin/$(APP_NAME)

