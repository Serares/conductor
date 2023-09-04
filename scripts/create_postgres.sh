#!/bin/bash

DOCKER=podman
GOOSE=goose
ROOT_DIR=$(pwd)

if command -v docker &>/dev/null; then
    printf "Docker looks to be available $(which docker)"
    DOCKER=docker
elif command -v podman &>/dev/null; then
    printf "Podman will be used. $(which podman)"
    DOCKER=podman
else
    printf "Docker and Podman not found on the machine"
    exit 1
fi

$DOCKER run -p 5432:5432 -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD -e POSTGRES_DB=$POSTGRES_DB -d --name=$POSTGRES_USER postgres:16rc1-alpine

printf "\nWaiting for PostgreSQL to be fully available."
until $DOCKER exec postgres pg_isready >/dev/null 2>&1; do
    printf "."
    sleep 5
done

printf "\nTrying to run migrations..."

if command -v $GOOSE &>/dev/null; then
    cd $ROOT_DIR/repository/migrations

    printf "\nRunning migrations $(pwd)"
    $GOOSE postgres "user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=disable" up
fi
