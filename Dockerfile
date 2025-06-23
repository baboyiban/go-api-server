# 빌드 단계
FROM golang:1.24-bookworm AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

FROM ubuntu:24.04

WORKDIR /app

RUN apt-get update && \
    apt-get install -y ca-certificates tzdata && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /main /app/main

CMD ["/app/main"]
