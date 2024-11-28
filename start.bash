#!/bin/bash

# just to make it easier as example
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5433
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=postgres
export POSTGRES_NAME=go-to-work

CONTAINER_NAME=go_to_work_db

if [ "$(docker inspect -f '{{.State.Running}}' $CONTAINER_NAME 2>/dev/null)" != "true" ]; then
  echo "Starting Postgres Container..."
  docker compose up -d $CONTAINER_NAME
else
  echo "Postgres Container Already Up"
fi

go run cmd/api/main.go