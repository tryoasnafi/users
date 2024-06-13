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

network:
	docker network create users-service-network

start-db:
	docker run --name postgresdb \
		-p 5432:5432 \
		--network users-service-network \
		-e POSTGRES_USER=ayousers \
		-e POSTGRES_PASSWORD=some@Password1221 \
		-e POSTGRES_DB=usersdb \
		-v users-service_postgres_data:/var/lib/postgresql/data \
		-d postgres:16

start-app:
	docker run --name users-service \
        -p 9090:9090 \
        --network users-service-network \
        --env-file .env \
        user-service:latest

migrate-db:
	@go run cmd/migration/main.go