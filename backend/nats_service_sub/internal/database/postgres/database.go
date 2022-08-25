package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"log"
)

const (
	user     = "cloneyd"
	password = "admin"
	host     = "database"
	port     = "5432"
	name     = "l0_task"
)

type Database struct {
	db *pgx.Conn
}

func NewConn() *Database {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, name)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalln(err)
	}

	return &Database{db: conn}
}

func (db *Database) Close() error {
	return db.db.Close(context.Background())
}

func (db *Database) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return db.db.Query(ctx, sql, args...)
}

func (db *Database) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.db.Exec(ctx, sql, args...)
}
