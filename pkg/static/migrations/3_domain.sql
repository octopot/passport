-- +migrate Up

CREATE DOMAIN "fingerprint" AS TEXT CHECK (VALUE ~ '^[^_\W]{32}$');

CREATE TABLE "control"
(
  "id"          BIGSERIAL PRIMARY KEY,
  "session"     UUID        NOT NULL,
  "fingerprint" FINGERPRINT NOT NULL,
  "context"     JSONB       NOT NULL,
  "created_at"  TIMESTAMP   NOT NULL DEFAULT now(),
  "updated_at"  TIMESTAMP   NULL     DEFAULT NULL,
  "deleted_at"  TIMESTAMP   NULL     DEFAULT NULL
);

CREATE DOMAIN "gaid" AS TEXT CHECK (VALUE ~ '^UA-\d+-\d+$');

CREATE TABLE "ga-beacon"
(
  "id"               BIGSERIAL PRIMARY KEY,
  "account_id"       UUID      NOT NULL,
  "google_analytics" GAID      NOT NULL,
  "created_at"       TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at"       TIMESTAMP NULL     DEFAULT NULL,
  "deleted_at"       TIMESTAMP NULL     DEFAULT NULL,
  UNIQUE ("account_id", "google_analytics")
);

-- +migrate Down

DROP TABLE "ga-beacon";

DROP DOMAIN "gaid";

DROP TABLE "control";

DROP DOMAIN "fingerprint";
