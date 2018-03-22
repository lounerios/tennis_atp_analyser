package main

import (
    "os"
    "fmt"
    "bufio"
    "encoding/csv"
    "io"
    "strconv"
    "gopkg.in/mgo.v2"

)

type player struct  {
     id int
     firstName string
     lastName string
     birthDate string
     country string
}

func NewPlayer(csv_line []string) *player {
    
    id, err := strconv.Atoi(csv_line[0])
    
    if (err != nil)  {
        return nil
    } 

    
    return &player{id, csv_line[1], csv_line[2], csv_line[3], csv_line[4]} 
}

func (p player) print() {
     fmt.Printf("%d %s\n", p.id, p.lastName)
}


type database struct {
    url string
    session mgo.Session
    db *mgo.Database
   
}

func (dbConn database) connect() error {
    session, err :=  mgo.Dial(dbConn.url)
    if (err != nil) {
        return err
    }
    
    dbConn.session = *session
    dbConn.db = session.DB("tennisatp")
    
    return nil
}

func (dbConn database) insertPlayer(p player) {
    c := dbConn.db.C("players")
    err := c.Insert(p)

    if (err != nil) {
       return
    }
}

func (dbConn database) close()  {
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

         p.print()
     }

    
}

