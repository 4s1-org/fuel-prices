package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Parse(start, end time.Time, c *Configuration) {
	loopOverFiles(start, end, c)
}

func loopOverFiles(start, end time.Time, c *Configuration) {
	numJobs := int(end.Sub(start).Hours() / 24)
	if numJobs <= 0 {
		return
	}

	jobs := make(chan time.Time, numJobs)
	results := make(chan time.Time, numJobs)

	// https://gobyexample.com/worker-pools
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results, c)
	}

	for date := start; date.After(end) == false; date = date.AddDate(0, 0, 1) {
		jobs <- date
	}
	close(jobs)

	// Warten, damit alle Daten durch sind
	for i := 1; i <= numJobs; i++ {
		<-results
	}
}

func worker(id int, jobs <-chan time.Time, results chan<- time.Time, c *Configuration) {
	for date := range jobs {
		parseFile(date, c)
		results <- date
	}
}

func parseFile(date time.Time, c *Configuration) {
	year := fmt.Sprintf("%04d", date.Year())
	month := fmt.Sprintf("%02d", date.Month())
	day := fmt.Sprintf("%02d", date.Day())

	fmt.Println(fmt.Sprintf("%s-%s-%s", year, month, day))

	srcPath := filepath.Join(c.TankerkoenigDataFolder, "prices", year, month)
	srcFilename := fmt.Sprintf("%s-%s-%s-prices.csv", year, month, day)
	srcFile2 := filepath.Join(srcPath, srcFilename)

	if _, err := os.Stat(srcFile2); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Sourcefile \"" + srcFile2 + "\" does not exists.")
		return
	}

	destPath := filepath.Join(c.CsvDataFolder, year, month)
	destFilename := fmt.Sprintf("%s-%s-%s.csv", year, month, day)
	destFile2 := filepath.Join(destPath, destFilename)

	srcFile, err := os.Open(srcFile2)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	err2 := os.MkdirAll(destPath, os.ModePerm)
	if err2 != nil {
		log.Fatal(err2)
	}

	destFile, err := os.Create(destFile2)
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	csvReader := csv.NewReader(srcFile)
	csvWriter := csv.NewWriter(destFile)
	defer csvWriter.Flush()
	firstRow := true

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if firstRow {
			header := []string{"date", "brand", "city", "street", "fueltype", "price"}
			csvWriter.Write(header)
			firstRow = false
			continue
		}

		timestampParsed, err := time.Parse(time.RFC3339, row[0][0:10]+"T"+row[0][11:22]+":00")
		if err != nil {
			log.Fatal(err)
		}
		timestampString := timestampParsed.Format(time.RFC3339)
		uuid := row[1]

		// ToDo: Das geht besser in Go
		for i := 0; i < len(c.Stations); i++ {
			station := c.Stations[i]
			if uuid == station.Id {
				details := []string{timestampString, station.Brand, station.City, station.Street}
				parseLine(csvWriter, row, details)
			}
		}
	}

	// ToDo: Add Brennpaste
}

func parseLine(csvWriter *csv.Writer, row []string, details []string) {
	diesel := row[2]
	e5 := row[3]
	e10 := row[4]
	dieselchange := row[5]
	e5change := row[6]
	e10change := row[7]

	// 0=keine Änderung, 1=Änderung, 2=Entfernt, 3=Neu
	// ToDo -1 und 2 beachten

	if dieselchange == "1" {
		csvWriter.Write(append(details, "Diesel", diesel))
	}
	if e5change == "1" {
		csvWriter.Write(append(details, "E5", e5))
	}
	if e10change == "1" {
		csvWriter.Write(append(details, "E10", e10))
	}
}
