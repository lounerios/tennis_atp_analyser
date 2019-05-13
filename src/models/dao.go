package models

import (
	"errors"
	"log"
)

const DB_SESSION_IS_NULL string = "DB session is NULL"

func InsertPlayer(dbConn *Database, p Player) error {
	if dbConn.db == nil {
		return errors.New(DB_SESSION_IS_NULL)
	}

	var op Player

	dbConn.db.Where(&Player{ID: p.ID}).First(&op)

	if op.ID != p.ID {
		dbConn.db.Create(&p)
	} else {
		return errors.New("Duplicate")
	}

	return nil
}

func InsertAtpRanking(dbConn *Database, ar AtpRanking) error {
	if dbConn.db == nil {
		return errors.New(DB_SESSION_IS_NULL)
	}

	var oar AtpRanking

	dbConn.db.Where(&AtpRanking{PlayerId: ar.PlayerId, Date: ar.Date}).First(&oar)

	if oar.PlayerId != ar.PlayerId {
		dbConn.db.Create(&ar)
	} else {
		return errors.New("Duplicate")
	}

	return nil
}

func InsertTournament(dbConn *Database, t Tournament) error {
	if dbConn.db == nil {
		return errors.New(DB_SESSION_IS_NULL)
	}

	var ot Tournament

	dbConn.db.Where(&Tournament{ID: t.ID}).First(&ot)
	log.Println(ot, t.ID)
	if ot.ID != t.ID {
		dbConn.db.Create(&t)
	} else {
		return errors.New("Duplicate")
	}

	return nil
}

func InsertMatch(dbConn *Database, m Match) error {
	if dbConn.db == nil {
		return errors.New(DB_SESSION_IS_NULL)
	}

	var om Match

	dbConn.db.First(&om, m.ID)

	if om.ID != m.ID {
		dbConn.db.Debug().Create(&m)

		for _, s := range m.Sets {
			s.MatchId = m.ID
			dbConn.db.Debug().Create(&s)
		}
	} else {
		return errors.New("Duplicate")
	}

	return nil

}
