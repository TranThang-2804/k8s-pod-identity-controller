VERSION 0.8
FROM golang:latest
LABEL maintainer="Tommy Tran Duc Thang <tranthang.dev@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

ci:
    ARG IMAGE_NAME=k8s-pod-identity-controller
    ARG TAG
    RUN echo "Starting CI..."
    BUILD +test
    RUN echo "Starting Building..."
    BUILD --pass-args +build

lint:
    FROM golangci/golangci-lint:latest
    COPY . .
    RUN echo "Starting Linting..."
    CMD ["golangci-lint", "run"] 


test:
    RUN echo "Starting Testing..."
    COPY . .
    RUN go test ./...

build:
    RUN echo "Starting Building..."
    ARG IMAGE_NAME
    ARG TAG
    COPY . .
    RUN go build -o main .
    EXPOSE 8080
    CMD ["./main"]
    SAVE IMAGE $IMAGE_NAME:$TAG
