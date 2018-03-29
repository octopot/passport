package domain

import "database/sql"

// Fingerprint represents a fingerprint of a user.
type Fingerprint struct {
	ID        int64
	Marker    string
	Value     string
	Counter   int
	CreatedAt string
	UpdatedAt sql.NullString
}
