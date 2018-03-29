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

     if (len(args) != 1) {
         fmt.Println("Please use the command:atp_read <filename>")
         return;
     }

     filename := args[0]

     fmt.Println("Reading file ", filename)

     file, err := os.Open(filename)

     utils.CheckErr(err)

     defer file.Close()

     dbConn := models.NewDatabase("mongodb://127.0.0.1")

     err = dbConn.Connect()
     utils.CheckErr(err)

     defer dbConn.Close()

     reader := csv.NewReader(bufio.NewReader(file))

     count := 0
     for {
          csv_line, err := reader.Read()
          if (err == io.EOF) {
              break;
          }

          count = count + 1
          fmt.Println(csv_line)
          if (count == 1) {
              continue
          }
          t := models.NewTournament(csv_line)
          models.InsertTournament(dbConn, *t)

          m := models.NewMatch(csv_line)
          models.InsertMatch(dbConn, *m)
     }
}
