
FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/main ./cmd/api/main.go

FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata && \
    update-ca-certificates

RUN adduser -D appuser
USER appuser

COPY --from=builder --chown=appuser:appuser /app/main /app/main

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/app/main"]