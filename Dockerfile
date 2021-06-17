FROM golang:alpine AS builder

WORKDIR /app
COPY . .
ENV DB_URL postgres://postgres:postgres@localhost:5432/mydb?sslmode=verify-ca&pool_max_conns=10


RUN go build -o build/kraken-meta-order cmd/kraken-meta-order/main.go

FROM alpine:3
WORKDIR /root
COPY --from=builder /app/build/kraken-meta-order .

EXPOSE 8080

CMD ["/root/kraken-meta-order"]