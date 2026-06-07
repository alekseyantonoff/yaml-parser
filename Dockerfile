# Build stage
FROM golang:alpine3.23 AS builder

WORKDIR /app

RUN apk update; \
    apk add --no-cache git curl bash; \
    apk upgrade --no-cache; \
    rm -rf /var/cache/apk/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o yaml-parser ./cmd/yaml-parser

# Runtime stage - using scratch (empty image) for smallest size
FROM scratch

WORKDIR /app

# Copy binary and timezone data
COPY --from=builder /app/yaml-parser .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./yaml-parser"]
CMD ["-help"]