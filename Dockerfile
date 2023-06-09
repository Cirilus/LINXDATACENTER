FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

ENTRYPOINT ["/app/main"]