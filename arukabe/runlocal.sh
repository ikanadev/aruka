#!/bin/bash

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

export DATABASE="postgres://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_DB}?sslmode=disable"
export MIGRATIONS_SOURCE="file://./migrations"
export GRPC_PORT=5000

go run ./cmd/grpc/main.go
