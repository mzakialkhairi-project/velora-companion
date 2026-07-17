#!/bin/bash
# Reset development environment

set -e

echo "Resetting development environment..."

# Stop and remove containers
docker compose down -v

# Remove build artifacts
rm -rf backend/bin
rm -rf mobile/build

echo "Reset complete!"
