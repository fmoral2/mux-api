# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

ARG APP_NAME
ENV APP_NAME ${APP_NAME}

ENV PATH="/opt/go/bin:${PATH}"
WORKDIR /opt/go/bin/

COPY . .
COPY ./resources/config*.yml ./resources/
RUN apk add --update alpine-sdk
RUN go mod tidy
RUN go build -o main github.com/morlfm/rest-api
RUN chmod +x /GO
CMD ["./main"]