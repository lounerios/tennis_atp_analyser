package models

import (
  "errors"
  "gopkg.in/mgo.v2/bson"
)

const DB_SESSION_IS_NULL string = "DB session is NULL"

func InsertPlayer(dbConn *Database, p Player) error{
    if (dbConn.Session == nil) {
        return errors.New(DB_SESSION_IS_NULL)
    }

    coll := dbConn.Db.C("players")
    err := coll.Insert(&p)

    return err
}

func InsertMatch(dbConn *Database, m Match) error{
    if (dbConn.Session == nil) {
        return errors.New("DB_SESSION_IS_NULL")
    }

    coll := dbConn.Db.C("matches")
    err := coll.Insert(&m)

    return err
}

func InsertTournament(dbConn *Database, t Tournament) error{
    if (dbConn.Session == nil) {
        return errors.New("DB_SESSION_IS_NULL")
    }

    coll := dbConn.Db.C("tournaments")

    n, err := coll.Find(bson.M{"id" : t.Id}).Count()

    if (err != nil) {
        return err
    }

    if ( n > 0) {
        return nil
    }

    err = coll.Insert(&t)

    return err
}
