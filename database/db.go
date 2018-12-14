package database

import (
	"database/sql"
	_ "github.com/lib/pq"	// package needed (for registering its drivers with the database/sql package) but never directly referenced in code
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"

)

type Config struct {
	Host string
	Port uint16
	User string
	Password string
	Database string
	SSLMode string
}

type Connection struct {
	// Db holds a sql.DB pointer that represents a pool of zero or more underlying connections - safe for concurrent use
	// by multiple goroutines -, with freeing/creation of new connections all managed by `sql/database` package.
	Db  *sql.DB
	cfg Config
}

// New returns a Connection with the sql.DB set with the postgres DB connection string in the configuration
func New(cfg Config) (connection Connection, err error) {
	if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" ||
		cfg.Password == "" || cfg.Database == "" {
		err = errors.Errorf(
			"All fields must be set (%s)",
			spew.Sdump(cfg))
		return
	}

	connection.cfg = cfg

	// The first argument corresponds to the driver name that the driver (in this case, `lib/pq`) used to register itself in `database/sql`.
	// The next argument specifies the parameters to be used in the connection.
	// Details about this string can be seen at https://godoc.org/github.com/lib/pq
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port, cfg.SSLMode)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't open connection to postgres database (%s)",
			spew.Sdump(cfg))
		return
	}

	fmt.Printf("Successfully connected to postgres database %s\n", cfg.Database)

	// Ping verifies if the connection to the database is alive or if a
	// new connection can be made.
	if err = db.Ping(); err != nil {
		err = errors.Wrapf(err,
			"Couldn't ping postgres database (%s)",
			spew.Sdump(cfg))
		return
	}

	connection.Db = db
	return
}

// Close performs the release of any resources that `sql/database` DB pool created. This is usually meant to be used in
// the exiting of a program or `panic`ing.
func (c *Connection) Close() (err error) {
	if c.Db == nil {
		return
	}

	if err = c.Db.Close(); err != nil {
		err = errors.Wrapf(err,
			"Error closing database connection",
			spew.Sdump(c.cfg))
	}

	return
}
