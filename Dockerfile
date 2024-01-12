# Stage 1: Build the Go binary
FROM golang:latest AS builder

LABEL maintainer="StefanWellhoner <stefanwellhoner@ymail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make build

# Stage 2: Create a minimal image to run the Go binary
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ .

EXPOSE 8080

CMD ["./task-forge"]