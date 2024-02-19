# Stage 1: Build the Go binary
FROM golang:latest AS builder

RUN apt-get update && apt-get install -y \
    && rm -rf /var/lib/apt/lists/*

LABEL maintainer="StefanWellhoner <stefanwellhoner@ymail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o task-forge main.go

# Stage 2: Create a minimal image to run the Go binary
FROM amd64/debian:stable-slim

WORKDIR /app

COPY --from=builder /app/ .

EXPOSE 8080

CMD ["./task-forge"]