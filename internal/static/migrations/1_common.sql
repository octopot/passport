-- +migrate Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "action" AS ENUM ('create', 'update', 'delete', 'restore');

-- +migrate Down

DROP TYPE "action";
