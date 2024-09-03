FROM golang:1.22.5 AS builder

WORKDIR /app

ARG SERVICE_NAME

COPY shared/middlewares /app/shared/middlewares
COPY shared/utils /app/shared/utils
COPY shared/config /app/shared/config

COPY ${SERVICE_NAME}/ /app/${SERVICE_NAME}/

WORKDIR /app/${SERVICE_NAME}

RUN go mod tidy

RUN go build -o /app/${SERVICE_NAME}

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/${SERVICE_NAME} /app/
CMD ["./${SERVICE_NAME}"]
