#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE db;
    GRANT ALL PRIVILEGES ON DATABASE unleash TO $POSTGRES_USER;
    CREATE DATABASE placio;
    GRANT ALL PRIVILEGES ON DATABASE placio TO $POSTGRES_USER;
EOSQL