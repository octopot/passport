-- +migrate Up

CREATE TABLE "account"
(
  "id"         UUID      NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name"       TEXT      NOT NULL,
  "created_at" TIMESTAMP NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP NULL                 DEFAULT NULL,
  "deleted_at" TIMESTAMP NULL                 DEFAULT NULL
);

CREATE TABLE "user"
(
  "id"         UUID      NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "account_id" UUID      NOT NULL,
  "name"       TEXT      NOT NULL,
  "created_at" TIMESTAMP NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP NULL                 DEFAULT NULL,
  "deleted_at" TIMESTAMP NULL                 DEFAULT NULL
);

CREATE TABLE "token"
(
  "id"         UUID      NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id"    UUID      NOT NULL,
  "expired_at" TIMESTAMP NULL                 DEFAULT NULL,
  "created_at" TIMESTAMP NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP NULL                 DEFAULT NULL,
  "deleted_at" TIMESTAMP NULL                 DEFAULT NULL
);

-- +migrate Down

DROP TABLE "token";

DROP TABLE "user";

DROP TABLE "account";
