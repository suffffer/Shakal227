FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/main.go

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/server .

CMD ["./server"]
