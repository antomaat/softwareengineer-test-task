package db

import (
        "log"

        "database/sql"
	_ "modernc.org/sqlite"

	"github.com/antomaat/softwareengineering-test-task/app/errors"
)

type Database struct {
    Conn *sql.DB
}

func NewDatabase(src string) (*Database, error) {
    conn, err := sql.Open("sqlite", src)
    if err != nil {
        log.Printf("Unable to open connection to database %v", err.Error())
        return nil, errortypes.InternalError 
    }
    if err := conn.Ping(); err != nil {
        log.Printf("Unable to open connection to database %v", err.Error())
        return nil, errortypes.InternalError 
    }

    return &Database{conn}, nil
}

func (db *Database) Close() error {
    return db.Conn.Close()
}

