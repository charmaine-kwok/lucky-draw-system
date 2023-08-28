# Use an official Golang runtime as a parent image
FROM golang:1.21.0-alpine3.18 AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Download swag CLI and initialise Swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12 && swag init

# # Build the Go app
RUN go mod download && go build -o main .

# Use an official Alpine Linux runtime as a parent image
FROM alpine:3.17.3

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the build stage to the new container
COPY --from=builder /app/main .

# Set the Gin mode to be release
ENV GIN_MODE=release

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the Go app when the container starts
CMD ["./main"]

