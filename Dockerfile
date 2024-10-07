# Use the official Golang image as the base image
FROM golang:1.23.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /aws-client-monitor

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY main.go .
COPY templates .

# Build the Go application
RUN go build -o aws-client-monitor .

# Create a smaller final image from scratch or alpine
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /aws-client-monitor

# Copy the compiled Go binary from the builder image
COPY --from=builder /aws-client-monitor/aws-client-monitor .
COPY templates /aws-client-monitor/templates

# Expose both UDP and TCP ports
EXPOSE 31000/udp
EXPOSE 8080/tcp

# Run the Go application
CMD ["./aws-client-monitor"]
