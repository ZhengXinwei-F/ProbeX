# Stage 1: Build the binary
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /probex

# Copy the source code into the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o probex .

# Stage 2: Create the minimal production image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder image
COPY --from=builder /probex/probex .

# Command to run the binary
CMD ["./probex"]
