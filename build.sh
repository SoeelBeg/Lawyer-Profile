#!/bin/sh

# Build the Go application
go build -o main cmd/main/main.go

# Make the binary executable
chmod +x main
