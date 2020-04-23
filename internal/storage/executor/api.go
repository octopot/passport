package executor

import "fmt"

const (
	postgresDialect = "postgres"
	mysqlDialect    = "mysql"
)

// New returns configured instance of abstract executor.
func New(dialect string) *Executor {
	exec := &Executor{dialect: dialect}
	switch exec.dialect {
	case postgresDialect:
	case mysqlDialect:
		fallthrough
	default:
		panic(fmt.Sprintf("not supported dialect %q is provided", exec.dialect))
	}
	return exec
}

// Executor provides access to abstract storage engine layers.
type Executor struct {
	dialect string
	factory struct{}
}

// Dialect returns executor dialect.
func (e *Executor) Dialect() string {
	return e.dialect
}
