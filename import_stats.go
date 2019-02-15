package main

import (
  "os"
  "fmt"
  "bufio"
  "encoding/csv"
  "io"
  "./src/utils"
  "./src/models"
  "flag"
  "time"
  "path/filepath"
  "strings"
)

var startTime = time.Now()

func getDuration() time.Duration{
  currentTime := time.Now()
  elapsed := currentTime.Sub(startTime)
  startTime = time.Now()

  return elapsed
}

func importFile(filename string, object string, dbConn *models.Database) error {
  fmt.Println("Import file:", filename, " at ", startTime)
  file, err := os.Open(filename)

  if (err != nil) {
    return err
  }
  defer file.Close()
  reader := csv.NewReader(bufio.NewReader(file))
  count := 0
  startTime = time.Now()


  for {
       csv_line, err := reader.Read()
       if (err == io.EOF) {
           break
       }

       if (count % 1000 == 0) {
           elapsed := getDuration()
           fmt.Println("Time for 1000 records:", elapsed, " records:", count)
       }

       switch object {
         case "ranking":
           r := models.NewAtpRanking(csv_line)
           err = models.InsertAtpRanking(dbConn, *r)
           count = count + 1
           r = nil

         case "player":
           p := models.NewPlayer(csv_line)
           err = models.InsertPlayer(dbConn, *p)
           p = nil

         case "match":
           t := models.NewTournament(csv_line)
           m := models.NewMatch(csv_line)

           fmt.Println("Tournament:",t.Name, " Match:", m.MatchNum)
           count = count + 1

          //models.InsertTournament(dbConn, *t)


          //models.InsertMatch(dbConn, *m)
        }

        if (err != nil) {
            fmt.Println("Error:", err)
        }else {
            count = count + 1
        }
  }

  elapsed := getDuration()
  fmt.Println("Time:", elapsed, " records:", count)

  return err
}

func main() {
     args := os.Args[1:]

     if (len(args) < 3 && len(args) > 4) {
	     fmt.Println("Please use the command:import_stats -object=<object> -input=<input> <sqlitedb_file> <input_value>")
         return;
     }

     objectPtr := flag.String("object", "player",  "The type of the object which is going to be imported. Options are player, ranking, match")
     inputPtr := flag.String("input", "file", "Type of the third argument. Options are file, directory")
     flag.Parse()

     dbFilename := args[2]
     filename := args[3]

     fmt.Println("The file of database: ", dbFilename)

     filters := map[string]string{
       "player"  : "atp_players_",
       "match"   : "atp_matches_",
       "ranking" : "atp_rankings_",
     }

     dbConn := models.NewDatabase(dbFilename)

     err := dbConn.Connect()

     utils.CheckErr(err)

     defer dbConn.Close()

     if (*inputPtr == "file") {
       err = importFile(filename, *objectPtr, dbConn)

     } else if (*inputPtr == "directory") {
       fmt.Println("Directory:", *inputPtr)
       fmt.Println("Filter:", filters[*objectPtr])

       depth := 1

       err := filepath.Walk(filename, func(path string, info os.FileInfo, err error) error{
               if err != nil {
                 fmt.Println("Cannot access the path")
                 return err
               }

               if(info.IsDir() && depth > 1) {
                 return filepath.SkipDir
               }

               if(strings.Index(info.Name(),filters[*objectPtr]) == 0){
                  err = importFile(path, *objectPtr, dbConn)
                  fmt.Println(err)
               }


               depth = depth + 1

               return nil
       })

       if err != nil {
         fmt.Println("Error walking the path", err)
         return
       }
     }

}
