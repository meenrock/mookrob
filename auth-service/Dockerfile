# Use the official Go image as a parent image
FROM golang:1.21 AS build

WORKDIR /app/src

COPY auth-service/src .

WORKDIR /app/src

ARG DB_HOST
ARG DB_PORT
ARG DB_NAME
ARG DB_USERNAME
ARG DB_PASSWORD

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
    && migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=require&sslrootcert=/shared/labsuser.pem" -path ./resources/migrations up

WORKDIR /shared

COPY shared .

WORKDIR /app/src/main

RUN go build -o main .

CMD ["./main"]