# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

ARG APP_NAME
ENV APP_NAME ${APP_NAME}

ENV PATH="/opt/go/bin:${PATH}"
WORKDIR /opt/go/bin/

COPY .. .
COPY ./deploy/docker-compose*.yml ./deploy/
RUN apk add --update alpine-sdk
RUN go mod tidy
CMD ["./main"]