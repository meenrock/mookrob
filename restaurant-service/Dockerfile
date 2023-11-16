# Use the official Go image as a parent image
FROM golang:1.21 AS build

WORKDIR /app/src

COPY restaurant-service/src .

WORKDIR /shared

COPY shared .

WORKDIR /app/src/main

RUN go build -o main .

CMD ["./main"]