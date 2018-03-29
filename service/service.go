package service

// New returns a new instance of Passport service.
func New(dao Storage) *Passport {
	return &Passport{dao: dao}
}

// Passport is the primary application service.
type Passport struct {
	dao Storage
}
