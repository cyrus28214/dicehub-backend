#! /bin/bash
source .env
POSTGRES_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME"
psql $POSTGRES_URL $@