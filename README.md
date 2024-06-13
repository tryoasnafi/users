# Users Service

## Description
This is users details service
 
## Tech Stacks
- Golang
- PostgreSQL (Database)
- Docker

## Framework & Library
- Echo (HTTP Framework): https://github.com/labstack/echo
- GORM (ORM): https://github.com/go-gorm/gorm
- GoDotEnv (Configuration): https://github.com/joho/godotenv
- Go Playground Validator (Validation): https://github.com/go-playground/validator
- Swaggo (Swagger v2): https://github.com/swaggo/swag
- Air (Auto-reload development): https://github.com/air-verse/air

## API Specs
API Specs in `docs` folder or you can run the project and access http://localhost:9090/api/docs/index.html

### Rebuild documentation
Installing swaggo cli
```sh
go install github.com/swaggo/swag/cmd/swag@latest
```
Generate docs
```sh
swag fmt
swag init --parseDependency --parseInternal -g cmd/api/main.go
```

## Development
1. Clone the project
```sh
git clone https://github.com/tryoasnafi/users
```
2. Setting .env
```sh
cd users
cp -p .env.example .env
```
3. Install project dependencies
```sh
go get ./...
```
4. Install Air (auto-reload) and run docker compose
```sh
go install github.com/air-verse/air@latest
make watch
```
> If you don't have GNU Make, you can run `docker compose -f compose.dev.yaml up -d --build`

## Database Migration
Database migration using GORM Auto migrate, make sure to set `.env`
```sh
go run cmd/migration/main.go
```

## Testing

### Run Unit Tests
```sh
go test ./...
```

## Production
Build users service image
```sh
make build
```
> or run `docker build -f Dockerfile.production . -t user-service:latest`
