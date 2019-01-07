#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER "passport" WITH PASSWORD 'passport';
    CREATE DATABASE "passport" WITH OWNER "passport";
    \c "passport";
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL
