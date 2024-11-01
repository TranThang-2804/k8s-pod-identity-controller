VERSION 0.8
FROM docker:latest
RUN apk update && \
    apk add --no-cache \
    curl \
    git \
    bash \
    make \
    gcc \
    libc6-compat \
    musl-dev
RUN curl -LO https://golang.org/dl/go1.17.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz && \
    rm go1.17.6.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
WORKDIR /app
COPY . .

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
    RUN go test ./...

build:
    RUN echo "Starting Building..."
    FROM DOCKERFILE .
    ARG IMAGE_NAME
    ARG TAG
    SAVE IMAGE $IMAGE_NAME:$TAG
