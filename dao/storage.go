package dao

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/kamilsk/passport/dao/postgres"
	"github.com/kamilsk/passport/domain"
)

// Must returns a new instance of the Storage or panics if it cannot configure it.
func Must(configs ...Configurator) *Storage {
	instance, err := New(configs...)
	if err != nil {
		panic(err)
	}
	return instance
}

// New returns a new instance of the Storage or an error if it cannot configure it.
func New(configs ...Configurator) (*Storage, error) {
	instance := &Storage{}
	for _, configure := range configs {
		if err := configure(instance); err != nil {
			return nil, err
		}
	}
	return instance, nil
}

// Connection returns database connection Configurator.
func Connection(driver, dsn string, open, idle int) Configurator {
	return func(instance *Storage) error {
		var err error
		instance.conn, err = sql.Open(driver, dsn)
		instance.conn.SetMaxOpenConns(open)
		instance.conn.SetMaxIdleConns(idle)
		return err
	}
}

// Configurator defines a function which can use to configure the Storage.
type Configurator func(*Storage) error

// Storage is an implementation of Data Access Object.
type Storage struct {
	conn *sql.DB
}

// Connection returns current database connection.
func (l *Storage) Connection() *sql.DB {
	return l.conn
}

// Dialect returns supported database dialect.
func (l *Storage) Dialect() string {
	return postgres.Dialect()
}

// UUID returns a new generated unique identifier.
func (l *Storage) UUID() (domain.UUID, error) {
	return postgres.UUID(l.conn)
}

// TakeFingerprint takes a user fingerprint and stores it.
func (l *Storage) TakeFingerprint(fp domain.Fingerprint) (domain.Fingerprint, error) {
	return postgres.TakeFingerprint(l.conn, fp)
}
