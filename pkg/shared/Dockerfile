FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY pkg/shared/ /app/pkg/shared/

COPY internal/ /app/internal/

ARG SERVICE_NAME

COPY cmd/${SERVICE_NAME}/ /app/cmd/${SERVICE_NAME}/

WORKDIR /app/cmd/${SERVICE_NAME}

RUN go mod tidy
RUN go build -o /app/${SERVICE_NAME}

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/${SERVICE_NAME} /app/

CMD ["./${SERVICE_NAME}"]
