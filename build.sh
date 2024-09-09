#!/bin/sh

# Install dependencies
go mod tidy

# Build the project
go build -o main ./cmd/main
