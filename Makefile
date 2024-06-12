format:
	@go fmt ./...

doc:
	@swag fmt
	@swag init \
		--parseDependency \
		--parseInternal \
		-g cmd/api/main.go

watch:
	@docker compose -f compose.dev.yaml up --build

up:
	@docker compose -f compose.dev.yaml up -d --build

down:
	@docker compose -f compose.dev.yaml down -v

restart: down up

test:
	go test -v ./...

build:
	docker build -f Dockerfile.production . -t user-service:latest

start-db:
	docker run --name postgresdb \
           --env-file ./.env \
           -p 5432:5432 \
           -d postgres:16

start-app:
	docker run --name user-service_1 \
           --env-file ./.env \
		   --network=host \
		   -p 8080:9090 \
           -d user-service:latest

migrate-db:
	@go run cmd/migration/main.go