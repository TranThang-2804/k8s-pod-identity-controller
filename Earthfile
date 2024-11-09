VERSION 0.8
FROM golang:latest
LABEL maintainer="Tommy Tran Duc Thang <tranthang.dev@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./*.go ./
COPY ./pkg ./pkg

ci:
    FROM alpine:latest
    ARG IMAGE_NAME='k8s-pod-identity-controller'
    ARG TAG='latest'
    RUN echo "Starting CI..."
    BUILD +lint
    BUILD +test
    BUILD --pass-args +build

lint:
    FROM golangci/golangci-lint:latest
    RUN echo "Starting Linting..."
    COPY ./*.go ./
    COPY ./pkg ./pkg
    CMD ["golangci-lint", "run", "-v"] 

test:
    RUN echo "Starting Testing..."
    RUN go test ./...

build:
    RUN echo "Starting Building..."
    ARG IMAGE_NAME
    ARG TAG
    RUN go build -o main .
    EXPOSE 8080
    CMD ["./main"]
    SAVE IMAGE $IMAGE_NAME:$TAG
