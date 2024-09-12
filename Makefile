include .env

GOARCH := amd64
GOOS := linux
APP := "./bin/webapp"

all: test qa build

build: templ css
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APP)

dev:
	@make -j3 run templ-proxy css-watch

qa: lint fmt vet

lint:
	golangci-lint run

vet:
	go vet ./...

fmt:
	gofumpt -l -w .

tidy:
	go mod tidy

css:
	npx tailwindcss -i ./assets/input.css -o ./static/output.css --minify

css-watch:
	npx tailwindcss --watch -i ./assets/input.css -o ./static/output.css

templ:
	templ generate

templ-proxy:
	@templ generate --watch --proxy=http://localhost:${PORT}

test:
	go test ./...

run:
	@air

migrate:
	go run ./cmd/migrate up

rollback:
	go run ./cmd/migrate down

seed:
	go run ./cmd/seed
	
clean:
	go clean
	rm -f bin/*
	rm -f coverage*.out
