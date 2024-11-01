VERSION 0.8
FROM docker:latest
RUN apk update && \
    apk add --no-cache \
    curl \
    git \
    bash \
    make \
    gcc \
    musl-dev
RUN curl -LO https://golang.org/dl/go1.17.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz && \
    rm go1.17.6.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
WORKDIR /app
COPY . .

ci:
    RUN echo "Starting CI..."
    RUN echo "Starting Testing..."
    BUILD +test
    RUN echo "Starting Building..."
    BUILD +build-image

test:
    RUN go test ./...

build-image:
    FROM DOCKERFILE .
    SAVE IMAGE my-image:latest
