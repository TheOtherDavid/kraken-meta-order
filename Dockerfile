FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o build/kraken-meta-order cmd/kraken-meta-order/main.go

FROM alpine:3
WORKDIR /root
COPY --from=builder /app/build/kraken-meta-order .

EXPOSE 8080

CMD ["/root/kraken-meta-order"]