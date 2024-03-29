run:
	@go run main.go
build:
	@go build -o bin/main main.go
up: 
	@docker-compose up -d
down:
	@docker-compose down
dbup:
	@go run migrate/up/main.go
dbdown:
	@go run migrate/down/main.go
dbreset:
	@go run migrate/down/main.go
	@go run migrate/up/main.go