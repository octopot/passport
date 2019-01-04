-- +migrate Up

CREATE DOMAIN "fingerprint" AS TEXT
  CHECK (
      length(VALUE) >= 32 AND
      length(VALUE) <= 64
    );

CREATE TABLE "event"
(
  "id"          BIGSERIAL PRIMARY KEY,
  "session"     UUID        NOT NULL,
  "fingerprint" fingerprint NOT NULL,
  "created_at"  TIMESTAMP   NOT NULL DEFAULT now()
);

-- +migrate Down

DROP TABLE "event";

DROP DOMAIN "fingerprint";
