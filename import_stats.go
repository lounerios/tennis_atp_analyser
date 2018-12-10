package main

import (
  "os"
  "fmt"
  "bufio"
  "encoding/csv"
  "io"
  "utils"
  "models"
  "flag"
)

func main() {
     args := os.Args[1:]
     if (len(args) != 3) {
	     fmt.Println("Please use the command:import_stats -object=<object> <sqlitedb_file> <filename>")
         return;
     }
     
     objectPtr := flag.String("object", "player",  "The type of the object which is going to be imported. Options are player, ranking, match")
     flag.Parse()
     dbFilename := args[1]
     filename := args[2]

     fmt.Println("The file of database: ", dbFilename)
     fmt.Println("The file of the statistics: ", filename)

     file, err := os.Open(filename)
     utils.CheckErr(err)

     defer file.Close()

     dbConn := models.NewDatabase(dbFilename)

     err = dbConn.Connect()

     utils.CheckErr(err)

     defer dbConn.Close()

     reader := csv.NewReader(bufio.NewReader(file))

     for {
          csv_line, err := reader.Read()
          if (err == io.EOF) {
              break;
          }

         utils.CheckErr(err);
         switch *objectPtr { 
         case "ranking":
             r := models.NewAtpRanking(csv_line)
	     models.InsertAtpRanking(dbConn, *r)
	     r = nil
         case "player":
	     p := models.NewPlayer(csv_line)
             models.InsertPlayer(dbConn, *p)
             p = nil
         case "match":
             count = count + 1
             if (count == 1) {
                 continue
             }
             //t := models.NewTournament(csv_line)
             //models.InsertTournament(dbConn, *t)

             //m := models.NewMatch(csv_line)
             //models.InsertMatch(dbConn, *m)

	}

     }
}
