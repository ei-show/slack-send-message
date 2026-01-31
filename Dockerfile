# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install ca-certificates in the builder stage
RUN apk --no-cache add ca-certificates

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o slack-send-message .

# Runtime stage - use scratch since we have a static binary
FROM scratch

# Copy ca-certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary from builder
COPY --from=builder /app/slack-send-message /slack-send-message

# Set the entrypoint
ENTRYPOINT ["/slack-send-message"]
