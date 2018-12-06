package models

import (
       "database/sql"
       _ "github.com/mattn/go-sqlite3"
)

type Database struct {
    filename string
    db *sql.DB
}

func NewDatabase(dbFilename string) *Database {
	dbConn := new(Database)
	dbConn.filename = dbFilename

	return dbConn
}

func (dbConn *Database) Connect() error {
     conn, err := sql.Open("sqlite3", dbConn.filename)

     if err != nil {
        return err
     }

     dbConn.db = conn
     
     return nil
}

func (dbConn *Database) Close() {
      if dbConn.db == nil {
	      return
      }

      dbConn.db.Close()
}

