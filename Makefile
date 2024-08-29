GOARCH := amd64
GOOS := linux
APP := "./bin/myapp"

.PHONY: all clean build models

check-quality:
	make lint
	make fmt
	make vet

# Append || true below if blocking local developement
lint:
	golangci-lint run --enable-all

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
	@templ generate --watch --proxy=http://localhost:3000

build:
	make templ
	make css
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APP)

test:
	go test ./...

run:
	@air

dev:
	@make -j3 run templ-proxy css-watch

migrate:
	go run ./cmd/migrate up

rollback:
	go run ./cmd/migrate down

models:
	sqlboiler psql

seed:
	go run ./cmd/seed
	
clean:
	go clean
	rm -f bin/*
	rm -f coverage*.out

all:
	make test
	make build

