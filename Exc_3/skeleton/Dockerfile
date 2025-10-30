# todo

FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN bash ./scripts/build-application.sh


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ordersystem /app/ordersystem
CMD ["/app/ordersystem"] 

