FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" \
    -o server ./cmd/api

    
FROM alpine:3.22

RUN apk --no-cache add ca-certificates tzdata

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/sql ./sql

USER app

EXPOSE 8080

CMD ["./server"]