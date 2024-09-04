# Use the official Golang image as a build stage
FROM golang:1.23 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o main cmd/main/main.go

# Set the necessary permissions
RUN chmod +x ./main

# Use a minimal base image for the final image
FROM debian:bullseye-slim

# Copy the compiled binary from the build stage
COPY --from=build /app/main /main

# Expose the port the app runs on
EXPOSE 8080

# Set the entrypoint to the binary
ENTRYPOINT ["/main"]
