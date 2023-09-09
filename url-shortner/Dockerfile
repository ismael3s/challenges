FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add make
RUN make build
FROM alpine
WORKDIR /app
COPY --from=builder /app/short.io .
CMD ["./main"]