package postgres

import (
	"database/sql"

	"github.com/kamilsk/passport/domain"
	"github.com/kamilsk/passport/errors"
)

const dialect = "postgres"

const (
	upsertFingerprint = `INSERT INTO "fingerprint" ("marker", "value") VALUES ($1, $2)
ON CONFLICT ("marker", "value") DO UPDATE
SET "counter" =  "fingerprint"."counter" + 1, "updated_at" = now()
RETURNING "id", "counter", "created_at", "updated_at"`
)

// Dialect returns supported database dialect.
func Dialect() string {
	return dialect
}

// UUID returns a new generated unique identifier.
func UUID(db *sql.DB) (domain.UUID, error) {
	var id domain.UUID
	row := db.QueryRow(`SELECT uuid_generate_v4()`)
	if err := row.Scan(&id); err != nil {
		return id, errors.Database(errors.ServerErrorMessage, err, "trying to populate UUID")
	}
	return id, nil
}

// TakeFingerprint takes a user fingerprint and stores it.
func TakeFingerprint(db *sql.DB, fp domain.Fingerprint) (domain.Fingerprint, error) {
	err := db.QueryRow(upsertFingerprint, fp.Marker, fp.Value).Scan(&fp.ID, &fp.Counter, &fp.CreatedAt, &fp.UpdatedAt)
	if err != nil {
		return fp, errors.Database(errors.ServerErrorMessage, err,
			"trying to insert fingerprint %q with marker %q", fp.Value, fp.Marker)
	}
	return fp, nil
}
