FROM golang:1.20 AS builder

WORKDIR /app


ARG SERVICE_NAME

COPY ${SERVICE_NAME}/ ./

RUN go mod init

RUN go mod tidy
RUN go build -o ${SERVICE_NAME}


FROM alpine:3.17


WORKDIR /app


COPY --from=builder /app/${SERVICE_NAME} ./


CMD ["./${SERVICE_NAME}"]
