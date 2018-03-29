-- +migrate Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "fingerprint" (
  "id"         BIGSERIAL PRIMARY KEY,
  "marker"     UUID        NOT NULL,
  "value"      VARCHAR(64) NOT NULL,
  "counter"    BIGINT      NOT NULL DEFAULT 1,
  "created_at" TIMESTAMP   NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP   NULL     DEFAULT NULL,
  UNIQUE ("marker", "value")
);

-- +migrate Down

DROP TABLE "fingerprint";
