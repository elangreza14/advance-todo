#!make
include local.env

up-stack:
	docker-compose --env-file local.env up --build -d

up:
	docker-compose --env-file local.env up -d

down:
	docker-compose --env-file local.env down

run-local:
	MODE=local go run main.go
	
run-live:
	go run main.go

FILENAME?=file-name
migrate-create:
	migrate create -ext sql -tz Asia/Jakarta -dir ./migrations -format "20060102150405" create_${FILENAME}

migrate-alter:
	migrate create -ext sql -tz Asia/Jakarta -dir ./migrations -format "20060102150405" alter_${FILENAME}

migrate-down:
	migrate -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOSTNAME}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL}" -path ${POSTGRES_MIGRATION_FOLDER} down

.PHONY: up-stack up down run-local run-live migrate-create migrate-alter migrate-down