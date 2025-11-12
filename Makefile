APP_NAME=shortlink
GOFILES=$(shell find . -name "*.go" -not -path "./vendor/*")

.PHONY: tidy build run docker up down swagger dev

tidy:
	go mod tidy

build:
	go build -o bin/$(APP_NAME) ./cmd

run:
	go run ./cmd/main.go

docker:
	docker-compose up -d

down:
	docker-compose down

swagger:
	swag init --dir ./cmd,./internal/api,./internal/service,./internal/repo,./internal/config,./internal/model,./internal/pkg --generalInfo main.go --output docs --parseInternal

fmt:
	go fmt ./...

dev:
	make swagger
	make run
