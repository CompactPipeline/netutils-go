FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /netutils .

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
COPY --from=builder /netutils /usr/local/bin/netutils

ENTRYPOINT ["netutils"]