package domain

import "database/sql"

// Marker represents a marker to identify a user.
type Marker struct {
	ID          string
	Fingerprint string
	CreatedAt   string
	UpdatedAt   sql.NullString
}
