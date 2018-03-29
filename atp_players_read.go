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
