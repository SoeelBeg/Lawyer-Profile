# Use the official Golang image to create a build artifact
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main cmd/main/main.go

# Start a new stage from scratch
FROM scratch

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/main"]
