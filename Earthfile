VERSION 0.8
FROM docker:latest

ci:
    RUN echo "This is a CI step"
    BUILD +test
    BUILD +build-image

test:
    RUN echo "This is a test step"

build-image:
    FROM earthly/dind:alpine-3.19-docker-25.0.5-r0
    WITH DOCKER --pull hello-world
      RUN docker run hello-world
    END
