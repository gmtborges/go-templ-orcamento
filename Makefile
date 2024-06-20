GOARCH := amd64
GOOS := linux
APP := "./bin/myapp"

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
	go fmt ./...

tidy:
	go mod tidy

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APP)

css:
	npx tailwindcss -i ./assets/input.css -o ./static/output.css --minify

css-watch:
	npx tailwindcss --watch -i ./assets/input.css -o ./static/output.css

templ:
	templ generate

templ-proxy:
	templ generate --watch --proxy=http://127.0.0.1:3000

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

all:
	make check-quality
	make test
	make build

