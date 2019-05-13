package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	filename string
	db       *gorm.DB
}

func NewDatabase(dbFilename string) *Database {
	dbConn := new(Database)
	dbConn.filename = dbFilename

	return dbConn
}

func (dbConn *Database) Connect() error {
	db, err := gorm.Open("sqlite3", dbConn.filename)

	if err != nil {
		return err
	}

	if !db.HasTable(&Player{}) {
		db.Debug().AutoMigrate(&Player{}, &Match{}, &MatchSet{}, &PlayerStats{}, AtpRanking{}, Tournament{})
	}

	dbConn.db = db

	return nil
}

func (dbConn *Database) Close() {
	if dbConn.db == nil {
		return
	}

	dbConn.db.Close()
}
