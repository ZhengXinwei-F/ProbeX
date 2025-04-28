# Stage 1: Build the binary
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /probex

# Copy the source code
COPY . .

# Build for the target architecture
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -o probex .

# Stage 2: Create minimal production image
FROM alpine:latest

# Create a non-root user (UID/GID 1000 is the standard non-privileged user in Alpine)
RUN addgroup -g 1000 appuser && \
    adduser -u 1000 -G appuser -D appuser && \
    mkdir -p /probex && \
    chown -R 1000:1000 /probex

# Set working directory and permissions
WORKDIR /probex
COPY --from=builder --chown=1000:1000 /probex/probex .

# Switch to the non-privileged user
USER 1000

# Run the program
CMD ["./probex"]
