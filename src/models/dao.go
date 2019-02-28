package models

import (
  "errors"
)

const DB_SESSION_IS_NULL string = "DB session is NULL"

func InsertPlayer(dbConn *Database, p Player) error{
    if dbConn.db == nil {
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
    if dbConn.db == nil {
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
*/

func GetTournamentById(dbConn *Database, id string) (*Tournament, error) {
    if dbConn.db == nil {
        return nil, errors.New(DB_SESSION_IS_NULL)
    }

    stmt,err := dbConn.db.Prepare("SELECT * FROM tournament WHERE id = ?")

    if err != nil {
       return nil, err
    }

    defer stmt.Close()

    rows,err := stmt.Query(id)

    defer rows.Close()

    if rows.Next() {
      t := new(Tournament)
      //todo: copy row data to tournament object

      return t, nil
    }
    return nil, nil
}

func InsertTournament(dbConn *Database, t Tournament) error{
    if dbConn.db == nil {
        return errors.New(DB_SESSION_IS_NULL)
    }

    trnTournament,err := GetTournamentById(dbConn, t.Id)

    if trnTournament != nil {
       return nil
    }

    stmt, err := dbConn.db.Prepare("INSERT INTO tournament(id, name, surface, draw_size, level, date) values(?,?,?,?,?,?)")

    if err != nil {
       return err
    }

    defer stmt.Close()

    _, err = stmt.Exec(t.Id, t.Name, t.Surface, t.DrawSize, t.Level, t.Date)

    return err
}
