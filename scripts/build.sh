#!/bin/bash
# Build script for Velora backend

set -e

echo "Building Velora backend..."

cd backend

# Build the API server
go build -o bin/api ./cmd/api

echo "Build complete!"
