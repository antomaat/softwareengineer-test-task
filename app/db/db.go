package db

import (
        "database/sql"
	_ "modernc.org/sqlite"
)

type Database struct {
    Conn *sql.DB
}

func NewDatabase(src string) (*Database, error) {
    conn, err := sql.Open("sqlite", src)
    if err != nil {
        return nil, err
    }
    if err := conn.Ping(); err != nil {
        return nil, err
    }

    return &Database{conn}, nil
}

func (db *Database) Close() error {
    return db.Conn.Close()
}

