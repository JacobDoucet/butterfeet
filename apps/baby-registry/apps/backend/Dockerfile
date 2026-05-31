# Build stage — installs forge, generates code, builds binary
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates

# Install forge so we can regenerate from spec at build time.
RUN go install github.com/JacobDoucet/forge@v1.0.5

WORKDIR /src

# Copy the entire app (spec + go module + custom code)
COPY . .

# Regenerate from spec (so the image is self-contained from spec → binary)
RUN forge build

WORKDIR /src/apps/backend
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/server .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/affiliate-backfill ./cmd/affiliate-backfill

# Runtime
FROM alpine:3.21
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /out/server ./server
COPY --from=builder /out/affiliate-backfill ./affiliate-backfill
EXPOSE 8080
CMD ["./server"]
