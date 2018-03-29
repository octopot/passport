package service

import "github.com/kamilsk/passport/domain"

// Storage defines the behavior of Data Access Object.
type Storage interface {
	// UUID returns a new generated unique identifier.
	UUID() (domain.UUID, error)
	// TakeFingerprint takes a user fingerprint and stores it.
	TakeFingerprint(domain.Fingerprint) (domain.Fingerprint, error)
}
