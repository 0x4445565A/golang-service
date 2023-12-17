FROM golang:1.21-alpine3.19 AS builder

WORKDIR /app
COPY . .
RUN go build -o service /app/cmd/service/.

FROM alpine:3.19 AS app

COPY --from=builder /app/service .

EXPOSE 8080
CMD ["./service"]