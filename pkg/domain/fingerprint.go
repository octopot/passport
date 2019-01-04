package domain

import "regexp"

var fingerprint = regexp.MustCompile(`(?i:^[0-9A-F]{32,64}$)`)

// Fingerprint represents a fingerprint of a user.
type Fingerprint struct {
	ID        int64
	Session   UUID
	Marker    string
	CreatedAt string
}

// IsValid returns true if the Fingerprint contains valid value.
func (f Fingerprint) IsValid() bool {
	return f.Marker != "" && fingerprint.MatchString(f.Marker)
}
