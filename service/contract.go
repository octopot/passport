package service

import "github.com/kamilsk/passport/domain"

// Storage defines the behavior of Data Access Object.
type Storage interface {
	// Marker returns the Marker by provided ID.
	Marker(domain.UUID) (domain.Marker, error)
}
