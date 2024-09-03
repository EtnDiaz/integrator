FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN	go mod init

RUN go mod tidy

RUN go build -o api-gateway

CMD ["./api-gateway"]
