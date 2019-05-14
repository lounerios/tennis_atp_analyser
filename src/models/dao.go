package models

import (
	"errors"
)

var (
	ErrDBNotFound = errors.New("Database is NULL")
	ErrDuplicate  = errors.New("Duplicate entry")
)

func InsertPlayer(dbConn *Database, p Player) error {
	if dbConn.db == nil {
		return ErrDBNotFound
	}

	var op Player

	dbConn.db.Where(&Player{ID: p.ID}).First(&op)

	if op.ID != p.ID {
		dbConn.db.Create(&p)
	} else {
		return ErrDuplicate
	}

	return nil
}

func InsertAtpRanking(dbConn *Database, ar AtpRanking) error {
	if dbConn.db == nil {
		return ErrDBNotFound
	}

	var oar AtpRanking

	dbConn.db.Where(&AtpRanking{PlayerId: ar.PlayerId, Date: ar.Date}).First(&oar)

	if oar.PlayerId != ar.PlayerId {
		dbConn.db.Create(&ar)
	} else {
		return ErrDuplicate
	}

	return nil
}

func InsertTournament(dbConn *Database, t Tournament) error {
	if dbConn.db == nil {
		return ErrDBNotFound
	}
	var ot Tournament

	dbConn.db.Where(&Tournament{ID: t.ID}).First(&ot)
	if ot.ID != t.ID {
		dbConn.db.Create(&t)
	} else {
		return ErrDuplicate
	}
	return nil
}

func InsertMatch(dbConn *Database, m Match) error {
	if dbConn.db == nil {
		return ErrDBNotFound
	}

	var om Match

	dbConn.db.Where(&Match{TournamentId: m.TournamentId, MatchNum: m.MatchNum}).First(&om)

	if m.MatchNum != om.MatchNum {
		dbConn.db.Create(&m)

		for _, s := range m.Sets {
			s.MatchId = m.ID
			dbConn.db.Create(&s)
		}
	} else {
		return ErrDuplicate
	}

	return nil
}
