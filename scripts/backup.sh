#!/bin/bash
# Backup script

set -e

BACKUP_DIR="./backups"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

echo "Creating backup..."

mkdir -p "$BACKUP_DIR"

# Backup database
docker exec velora-postgres pg_dump -U postgres velora > "$BACKUP_DIR/db_$TIMESTAMP.sql"

echo "Backup complete: $BACKUP_DIR/db_$TIMESTAMP.sql"
