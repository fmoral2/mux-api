FROM golang:alpine

ARG APP_NAME
ARG AWS_REGION
ENV AWS_REGION ${AWS_REGION}
ENV APP_NAME ${APP_NAME}


WORKDIR /app


COPY . .

#ENV PATH="/opt/go/bin:${PATH}"


RUN apk add --update alpine-sdk
RUN go mod tidy
RUN go build -o main ./cmd/

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown -R appuser:appgroup /app && chmod -R 755 /app
USER appuser
 

CMD ["./main"]
#CMD ["sh", "-infra", "while true; do sleep 1000; done"]
#CMD ["docker-compose", "up", "./main"]