package main

import (
  "os"
  "fmt"
  "bufio"
  "encoding/csv"
  "io"
  "utils"
  "models"
)

func main() {
     args := os.Args[1:]
     fmt.Println(len(args))
     if (len(args) != 2) {
         fmt.Println("Please use the command:atp_players_import <sqlitedb_file> <filename>")
         return;
     }

     dbFilename := args[0]
     filename := args[1]

     fmt.Println("Reading file ", filename)

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

         p := models.NewPlayer(csv_line)
         models.InsertPlayer(dbConn, *p)
         p = nil
     }
}
