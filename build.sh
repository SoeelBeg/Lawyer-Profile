#!/bin/bash

# Build the Go binary
go build -o main cmd/main/main.go

# Set execute permissions
chmod +x main
