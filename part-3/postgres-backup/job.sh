#!/usr/bin/env bash
set -e

if [ $URL ]
then
  TIMESTAMP=$(date +%Y%m%d_%H%M%S)
  BACKUP_FILE="backup_${TIMESTAMP}.sql"
  
  echo "Creating backup..."
  pg_dump -v $URL > "/tmp/$BACKUP_FILE"
  
  echo "Uploading to Google Cloud Storage..."
  gsutil cp "/tmp/$BACKUP_FILE" "gs://$BUCKET_NAME/$BACKUP_FILE"
  
  echo "Cleaning up..."
  rm "/tmp/$BACKUP_FILE"
fi