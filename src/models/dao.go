package models

import (
  "errors"
)

const DB_SESSION_IS_NULL string = "DB session is NULL"

func InsertPlayer(dbConn *Database, p Player) error{
    if (dbConn.db == nil) {
        return errors.New(DB_SESSION_IS_NULL)
    }

    stmt, err := dbConn.db.Prepare("INSERT INTO player(id, firstname, lastname, status, birthdate, country) values(?,?,?,?,?,?)")
    
    if err != nil {
	    return err
    }

    defer stmt.Close()

    _, err = stmt.Exec(p.Id, p.Firstname, p.Lastname, p.Status, p.Birthdate, p.Country)   

    return err
}

func InsertAtpRanking(dbConn *Database, ar AtpRanking) error{
    if (dbConn.db == nil) {
        return errors.New(DB_SESSION_IS_NULL)
    }

    stmt, err := dbConn.db.Prepare("INSERT INTO atp_ranking(date, number, player_id, points) values(?,?,?,?)")

    if err != nil {
            return err
    }

    defer stmt.Close()

    _, err = stmt.Exec(ar.Date, ar.Number, ar.Player_Id, ar.Points)

    return err
}

/*
func InsertMatch(dbConn *Database, m Match) error{
    if (dbConn.Session == nil) {
        return errors.New("DB_SESSION_IS_NULL")
    }

    coll := dbConn.Db.C("match")
    err := coll.Insert(&m)

    return err
}

func InsertTournament(dbConn *Database, t Tournament) error{
    if (dbConn.Session == nil) {
        return errors.New("DB_SESSION_IS_NULL")
    }

    coll := dbConn.Db.C("tournament")

    n, err := coll.Find(bson.M{"id" : t.Id}).Count()

    if (err != nil) {
        return err
    }

    if ( n > 0) {
        return nil
    }

    err = coll.Insert(&t)

    return err
}*/
