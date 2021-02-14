# Build environment.

FROM golang:1.15.8-alpine3.13 AS builder

WORKDIR /app

COPY main.go .

RUN go build -o http-slow-response

# Runtime environment.

FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/http-slow-response .

EXPOSE 80

CMD ["/app/http-slow-response"]
