package domain

import (
	"database/sql"
	"regexp"
)

var fingerprint = regexp.MustCompile(`(?i:^[0-9A-F]{32,64}$)`)

// Fingerprint represents a fingerprint of a user.
type Fingerprint struct {
	ID        int64
	Marker    string
	Value     string
	Counter   int
	CreatedAt string
	UpdatedAt sql.NullString
}

// IsValid returns true if the Fingerprint contains valid value.
func (f Fingerprint) IsValid() bool {
	return f.Value != "" && fingerprint.MatchString(f.Value)
}
