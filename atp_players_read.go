package main

import (
    "os"
    "fmt"
    "bufio"
    "encoding/csv"
    "io"
    "gopkg.in/mgo.v2"
)

type player struct  {
     Id string            `json: "Id", bson:"id"`
     Firstname string     `json: "Firstname", bson:"Firstname"`
     Lastname string      `json: "Lastname", bson:"Lastname"`
     Status string        `json: "Status", bson:"Status"`
     Birthdate string     `json: "Birthdate", bson:"BirtDate"`
     Country string       `json: "Country", bson:"Country"`
}

func NewPlayer(csv_line []string) *player {
    return &player{csv_line[0], csv_line[1], csv_line[2], csv_line[3], csv_line[4], csv_line[5]} 
}

func (p player) print() {
     fmt.Printf("%d %s\n", p.Id, p.Lastname)
}


type database struct {
    url string
    session *mgo.Session
    db *mgo.Database
}

func (dbConn *database) connect() error {
    session, err :=  mgo.Dial(dbConn.url)
    if (err != nil) {
        return err
    }
    
    dbConn.session = session
    dbConn.db = dbConn.session.DB("tennisatp")
    dbConn.session.SetMode(mgo.Monotonic, true)

    fmt.Println("Connect to ", dbConn.url)
    return nil
}

func (dbConn *database) insertPlayer(p player) {
    if (dbConn.session == nil) {
        return
    }
 
    fmt.Println(p)
    coll := dbConn.db.C("players")
    err := coll.Insert(&p)

    if (err != nil) {
       checkErr(err)
       return
    }
}

func (dbConn database) close()  {
   if (dbConn.session == nil) {
       return
   }
   fmt.Println("Close connection")
   dbConn.session.Close()
}



func checkErr(e error) {
    if (e != nil) {
        panic(e)
    }
}

func main() {
     args := os.Args[1:]

     if (len(args) != 1) {
         fmt.Println("Please use the command:atp_read <filename>")
         return;
     }
  
     filename := args[0]

     fmt.Println("Reading file ", filename)
     
     file, err := os.Open(filename)
     
     checkErr(err)
      
     defer file.Close()

     dbConn := new(database)
     dbConn.url = "mongodb://127.0.0.1"

     err = dbConn.connect()
     checkErr(err)

     defer dbConn.close()

     reader := csv.NewReader(bufio.NewReader(file))
       
     for {
          csv_line, err := reader.Read()
          if (err == io.EOF) {
              break;
          }
           
         checkErr(err);
         
         p := NewPlayer(csv_line)
         dbConn.insertPlayer(*p)
         p = nil       
     }

    
}

