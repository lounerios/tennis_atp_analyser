package models

import (
    "gopkg.in/mgo.v2"
)

type Database struct {
    url string
    Session *mgo.Session
    Db *mgo.Database
}

func NewDatabase(db_url string) *Database {
  dbConn := new(Database)
  dbConn.url = db_url

  return dbConn
}

func (dbConn *Database) Connect() error {
    session, err :=  mgo.Dial(dbConn.url)
    if (err != nil) {
        return err
    }

    dbConn.Session = session
    dbConn.Db = dbConn.Session.DB("tennisatp")
    dbConn.Session.SetMode(mgo.Monotonic, true)

    return nil
}


func (dbConn Database) Close()  {
   if (dbConn.Session == nil) {
       return
   }

   dbConn.Session.Close()
}
