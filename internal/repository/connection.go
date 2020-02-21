package repo

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	// driver for postgres
	_ "github.com/lib/pq"
)

// Repository keeps all messages.
type Repository struct {
	db *sql.DB
}

// Configuration is setting for connection to database
type Configuration struct {
	DBName   string
	Address  string
	User     string
	Password string
}

// New creates new instance of database connection
func New(ctx context.Context, conf Configuration) (Repository, error) {
	r := Repository{}
	host, port, err := net.SplitHostPort(conf.Address)
	if err != nil {
		return r, err
	}

	err = r.connect(host, port, conf.User, conf.Password, conf.Password)
	if err != nil {
		return r, err
	}

	err = r.ping(ctx)
	if err != nil {
		return r, err
	}

	err = r.createTables(ctx)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (r *Repository) connect(host, port, user, password, dbName string) (err error) {
	r.db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName))
	return
}

func (r Repository) ping(ctx context.Context) error {
	return r.db.PingContext(ctx)
}

func (r Repository) createTables(ctx context.Context) (err error) {
	_, err = r.db.ExecContext(ctx,
		`
    CREATE TABLE IF NOT EXISTS messages (
      id          UUID          PRIMARY KEY,
      authorName  VARCHAR(50)   NOT NULL,
      created_at  TIMESTAMPTZ   NOT NULL,
      message     VARCHAR(300)  NOT NULL
    );
    `,
	)
	return
}

// Close connection
func (r Repository) Close() {
	r.db.Close()
}
