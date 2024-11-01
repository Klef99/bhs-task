include .env
export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

up: ### Run docker-compose
	docker-compose up --build -d && docker-compose logs -f
.PHONY: up

down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: down

dev-up: ### Up infrastructure
	docker-compose up -d postgres && docker-compose logs -f
.PHONY: dev-up

dev-down: ### Down infrastructure
	docker-compose down postgres --remove-orphans
.PHONY: dev-down

swag-v1: ### swag init
	swag init -g internal/controller/http/v1/router.go
.PHONY: swag-v1

mock-generate: ### generate mocks
	mockgen -source=internal/usecase/interfaces.go -destination=internal/usecase/mock_test.go -package=usecase_test 
.PHONY: mock-generate

rm-volume: ### remove docker volume
	docker volume rm bhs-task_pg-data
.PHONY: rm-volume

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test

cover: ### create cover file report
	go test -short -count=1 -race -coverprofile=coverage.out ./internal/...
	go tool cover -html coverage.out
	rm coverage.out
.PHONY: cover

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path migrations -database '$(PG_URL)?sslmode=disable' up
.PHONY: migrate-up

bin-deps:
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	GOBIN=$(LOCAL_BIN) go install go.uber.org/mock/mockgen@latest
