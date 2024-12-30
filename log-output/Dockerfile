# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /build

# Copy only dependency files first
COPY go.mod ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:3.18

LABEL maintainer="Project Maintainer"
LABEL description="Log Output Application"

WORKDIR /app

# Add non-root user
RUN adduser -D appuser

# Copy binary from builder
COPY --from=builder /build/main .

# Use non-root user
USER appuser

# Use exec form with full path
ENTRYPOINT ["/app/main"]