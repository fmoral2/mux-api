# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

ARG APP_NAME
ARG AWS_REGION
ENV AWS_REGION ${AWS_REGION}
ENV APP_NAME ${APP_NAME}

ENV PATH="/opt/go/bin:${PATH}"
WORKDIR /opt/go/bin/

COPY .. .

RUN apk add --update alpine-sdk
RUN go mod tidy
CMD ["docker-compose", "up", "./main"]
