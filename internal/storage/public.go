package storage

import (
	"context"

	"go.octolab.org/ecosystem/passport/internal/domain"
	"go.octolab.org/ecosystem/passport/internal/storage/executor/postgres"
)

// StoreFingerprint takes a user fingerprint and stores it.
func (storage *Storage) StoreFingerprint(ctx context.Context, fp domain.Fingerprint) (domain.Fingerprint, error) {
	var fingerprint domain.Fingerprint

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return fingerprint, connErr
	}
	defer func() { _ = closer() }()

	return postgres.StoreFingerprint(conn, ctx, fp)
}

// UUID returns a new generated unique identifier.
func (storage *Storage) UUID(ctx context.Context) (domain.UUID, error) {
	var uuid domain.UUID

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return uuid, connErr
	}
	defer func() { _ = closer() }()

	return postgres.UUID(conn, ctx)
}
