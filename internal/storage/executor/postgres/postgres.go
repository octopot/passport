package postgres

import (
	"context"
	"database/sql"

	"go.octolab.org/ecosystem/passport/internal/domain"
	"go.octolab.org/ecosystem/passport/internal/errors"
)

// StoreFingerprint takes a user fingerprint and stores it.
func StoreFingerprint(conn *sql.Conn, ctx context.Context, fp domain.Fingerprint) (domain.Fingerprint, error) {
	q := `INSERT INTO "event" ("session", "fingerprint") VALUES ($1, $2) RETURNING "id", "created_at"`
	err := conn.QueryRowContext(ctx, q, fp.Session, fp.Marker).Scan(&fp.ID, &fp.CreatedAt)
	if err != nil {
		return fp, errors.Database(errors.ServerErrorMessage, err,
			"trying to insert fingerprint %q with marker %q", fp.Marker, fp.Session)
	}
	return fp, nil
}

// UUID returns a new generated unique identifier.
func UUID(conn *sql.Conn, ctx context.Context) (domain.UUID, error) {
	var id domain.UUID
	q := `SELECT uuid_generate_v4()`
	row := conn.QueryRowContext(ctx, q)
	if err := row.Scan(&id); err != nil {
		return id, errors.Database(errors.ServerErrorMessage, err, "trying to populate UUID")
	}
	return id, nil
}
