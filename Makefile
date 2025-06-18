run:
	go run ./cmd/server

build:
	go build -o bin/app ./cmd/server

lint:
	golangci-lint run

format:
	gofmt -w . && goimports -w .

test:
	go test ./...

migrate:
	go run ./cmd/migrate