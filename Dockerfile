FROM golang:1.24.1-alpine AS builder
COPY . /go/src/rubik-scrumbler
WORKDIR /go/src/rubik-scrumbler
RUN go mod download
RUN go build -o ./bin/grpc_server cmd/grpc_server/main.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /go/src/rubik-scrumbler/bin/grpc_server .
COPY .env .
CMD ["./grpc_server"]
