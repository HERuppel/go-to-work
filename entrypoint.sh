#!/bin/bash

echo "Waiting for the database $POSTGRES_NAME..."
until nc -z -v -w30 $POSTGRES_HOST $POSTGRES_PORT; do
  echo "Waiting for the database on $POSTGRES_HOST:$POSTGRES_PORT..."
  sleep 1
done

echo "Running migrations up ..."
migrate -path=internal/database/migrations -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_NAME?sslmode=disable" -verbose up

echo "Initializing API..."
exec "$@"