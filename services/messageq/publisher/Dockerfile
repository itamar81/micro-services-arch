#docker buildx build . --tag ghcr.io/itamar81/micro-services-arch/messageq:1.1 --push --platform linux/amd64,linux/arm64
FROM golang:1.21.0 as builder
WORKDIR /app
# COPY go.mod ./

COPY . .
# COPY go.sum .
# RUN go mod download

# COPY *.go .
RUN go mod init  github.com/itamar81/micro-service-arch/messageq/publisher
RUN go get github.com/rabbitmq/amqp091-go
RUN go build -o messageq -tags netgo -a -v
FROM alpine:3.18.3
WORKDIR /app
COPY --from=builder /app/messageq .
ENTRYPOINT [ "./messageq" ]
