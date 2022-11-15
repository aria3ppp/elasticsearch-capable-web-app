ENVFILE ?= .env
$(if $(wildcard $(ENVFILE)), \
	$(foreach VAR, $(shell sed -ne 's/ *\#.*$$//; /./ s/=.*$$// p' $(ENVFILE)), \
		$(if $($(VAR)),, \
			$(eval $(shell \
					echo export $(VAR)=$(shell sed -nr 's/$(VAR)=(.+)/\1/p' $(ENVFILE)) \
			)) \
		) \
	) \
)

MIGRATE_DSN ?= "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable"
MIGRATE := docker run --rm -v $(shell pwd)/migrations:/migrations --user "$(shell id -u):$(shell id -g)" --network host migrate/migrate:4 -path=/migrations -database "$(MIGRATE_DSN)"

DOCKER_COMPOSE_DEPENDENCIES := docker compose -f docker-compose.dependencies.yml
DOCKER_COMPOSE_SERVER := $(DOCKER_COMPOSE_DEPENDENCIES) -f docker-compose.server.yml

.PHONY: default
default: help

# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: dependencies-up
dependencies-up: ## create and start dependency services
	$(DOCKER_COMPOSE_DEPENDENCIES) up -d

.PHONY: dependencies-down
dependencies-down: ## stop and remove dependency services
	$(DOCKER_COMPOSE_DEPENDENCIES) down

.PHONY: server-up
server-up: ## create and start server
	$(DOCKER_COMPOSE_SERVER) up -d

.PHONY: server-down
server-down: ## stop and remove server
	$(DOCKER_COMPOSE_SERVER) down

.PHONY: test
test: ## run unit tests
	go test -covermode=count -coverprofile=coverage.out -short ./models

# .PHONY: test-integration
# test-integration: ## run integration tests
# 	go test -covermode=count -coverprofile=coverage.out ./internal/repo/ -run "TestDBRepoSuite"

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage.out

# .PHONY: test-integration-cover
# test-integration-cover: test-integration ## run integration tests and show integration test coverage information
# 	go tool cover -html=coverage.out
	
.PHONY: test-arg-cover
test-arg-cover: ## run unit tests by passing $ARG env value to 'go test' command and show test coverge information
	go test -covermode=count -coverprofile=coverage.out $(ARG)
	go tool cover -html=coverage.out

.PHONY: run
run: ## run main package
	go run .

.PHONY: build
build: ## build main package
	CGO_ENABLED=0 go build -o server .

.PHONY: generate
generate: ## run 'go generate' for all packages
	go generate ./...

.PHONY: lint
lint: ## run staticcheck
	@staticcheck ./...

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migrate step..."
	@$(MIGRATE) down 1

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir /migrations/ $${name}

.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop -f
	@echo "Running all database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-arg
migrate-arg: ## run migration command with argument ARG
	@echo "Running migration command with argument: $(ARG)"
	@$(MIGRATE) $(ARG)
