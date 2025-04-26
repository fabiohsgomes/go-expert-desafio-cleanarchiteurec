FROM golang:1.24.2-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy \
&& cd cmd/ordersystem \
&& go build -o ordersystem main.go wire_gen.go

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/cmd/ordersystem/ordersystem ordersystem
COPY --from=builder /app/cmd/ordersystem/.env .env
CMD [ "./ordersystem" ]