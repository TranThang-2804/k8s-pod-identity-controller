VERSION 0.8
FROM golang:latest
LABEL maintainer="Tommy Tran Duc Thang <tranthang.dev@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

ci:
    ARG IMAGE_NAME
    ARG TAG
    RUN echo "Starting CI..."
    BUILD +test
    RUN echo "Starting Building..."
    BUILD --pass-args +build

lint:
    RUN echo "Starting Linting..."

test:
    RUN echo "Starting Testing..."
    COPY . .
    RUN go test ./...

build:
    RUN echo "Starting Building..."
    FROM DOCKERFILE .
    ARG IMAGE_NAME
    ARG TAG
    SAVE IMAGE $IMAGE_NAME:$TAG
