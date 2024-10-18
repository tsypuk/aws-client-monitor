# Use the official Golang image as the base image to vuild go back-end
FROM golang:1.23.1-alpine AS builder-go
# Set the working directory inside the container
WORKDIR /aws-client-monitor
# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy the rest of the application code
COPY cmd ./cmd
COPY docs ./docs
COPY internal ./internal
# Build the Go application
RUN go build -o aws-client-monitor ./cmd/main.go



# Stage 1: Build the React app
FROM node:22.4-alpine AS builder-react
# Set working directory inside the container
WORKDIR /app
COPY react-admin ./
RUN npm install --force
# Build the React app for production
RUN npm run build



# Create a smaller final image from alpine
FROM alpine:latest
# Set the working directory inside the container
WORKDIR /aws-client-monitor
# Copy the compiled Go binary from the builder image
COPY --from=builder-go /aws-client-monitor/aws-client-monitor .
COPY --from=builder-react /app/build ./react-admin/build
# Expose both UDP and TCP ports
EXPOSE 31000/udp
EXPOSE 8080/tcp
# Run the Go application
CMD ["./aws-client-monitor"]
