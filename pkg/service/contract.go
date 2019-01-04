package service

import (
	"context"

	"github.com/kamilsk/passport/pkg/domain"
)

// Storage defines the behavior of Data Access Object.
type Storage interface {
	// StoreFingerprint takes a user fingerprint and stores it.
	StoreFingerprint(context.Context, domain.Fingerprint) (domain.Fingerprint, error)
	// UUID returns a new generated unique identifier.
	UUID(context.Context) (domain.UUID, error)
}
