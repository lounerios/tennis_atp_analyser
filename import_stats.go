package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"./src/models"
	"./src/utils"
)

var (
	startTime time.Time
	beginTime = time.Now()
)

func getDuration() time.Duration {
	currentTime := time.Now()
	elapsed := currentTime.Sub(startTime)
	startTime = time.Now()

	return elapsed
}

func importFile(filename string, object string, dbConn *models.Database) error {

	log.Println("Import file:", filename, " at ", startTime)
	file, err := os.Open(filename)

	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))
	count := 0
	startTime = time.Now()

	for {

		csv_line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if count%1000 == 0 {
			elapsed := getDuration()
			log.Println("Time for 1000 records:", elapsed, " records:", count)
		}

		switch object {
		case "ranking":
			r := models.NewAtpRanking(csv_line)
			err = models.InsertAtpRanking(dbConn, *r)

		case "player":
			p := models.NewPlayer(csv_line)
			err = models.InsertPlayer(dbConn, *p)

		case "match":
			t := models.NewTournament(csv_line)
			m := models.NewMatch(csv_line)
			err = models.InsertTournament(dbConn, *t)
			if err == models.ErrDuplicate {
				continue
			}

			err = models.InsertMatch(dbConn, *m)
		}

		if err != nil && err != models.ErrDuplicate {
			log.Println(err)
		} else {
			count = count + 1
		}

	}

	elapsed := time.Now().Sub(beginTime)
	log.Println("Time:", elapsed, " records:", count)

	return nil
}

func main() {
	args := os.Args[1:]

	if len(args) < 3 || len(args) > 4 {
		fmt.Println("Please use the command:import_stats -object=<object> -input=<input> <sqlitedb_file> <input_value>")
		return
	}

	objectPtr := flag.String("object", "player", "The type of the object which is going to be imported. Options are player, ranking, match")
	inputPtr := flag.String("input", "file", "Type of the third argument. Options are file, directory")
	flag.Parse()

	dbFilename := args[2]
	filename := args[3]

	filters := map[string]string{
		"player":  "atp_players_",
		"match":   "atp_matches_",
		"ranking": "atp_rankings_",
	}

	dbConn := models.NewDatabase(dbFilename)

	err := dbConn.Connect()

	utils.CheckErr(err)

	defer dbConn.Close()

	if *inputPtr == "file" {
		err = importFile(filename, *objectPtr, dbConn)

	} else if *inputPtr == "directory" {
		log.Println("Directory:", *inputPtr)
		log.Println("Filter:", filters[*objectPtr])

		depth := 1

		err := filepath.Walk(filename, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println("Cannot access the path")
				return err
			}

			if info.IsDir() && depth > 1 {
				return nil
			}

			if strings.Index(info.Name(), filters[*objectPtr]) == 0 {
				err = importFile(path, *objectPtr, dbConn)
				if err != nil {
					return err
				}
			}

			depth = depth + 1

			return nil
		})

		if err != nil {
			log.Println("Error walking the path", err)
			return
		}
	}
}
