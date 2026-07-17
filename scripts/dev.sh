#!/bin/bash
# Development script for Velora

set -e

echo "Starting Velora development environment..."

# Start Docker services
docker compose up -d

# Wait for services to be ready
echo "Waiting for services to be ready..."
sleep 5

echo "Development environment is ready!"
echo "Backend: http://localhost:8080"
echo "MinIO Console: http://localhost:9001"
echo "Mailpit Web: http://localhost:8025"
