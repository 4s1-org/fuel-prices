package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

	fmt.Printf("%s-%s-%s", year, month, day)
	fmt.Println()

	srcPath := filepath.Join(c.TankerkoenigDataFolder, "prices", year, month)
	srcFilename := fmt.Sprintf("%s-%s-%s-prices.csv", year, month, day)
	srcFile2 := filepath.Join(srcPath, srcFilename)

	if _, err := os.Stat(srcFile2); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Sourcefile \"" + srcFile2 + "\" does not exists.")
		return
	}

	destPath := filepath.Join(c.CsvDataFolder, year, month)
	destFilename := fmt.Sprintf("%s-%s-%s.csv", year, month, day)
	destFilenameInfluxDb := fmt.Sprintf("%s-%s-%s-influxdb.txt", year, month, day)
	destFile2 := filepath.Join(destPath, destFilename)
	destFile2InfluxDb := filepath.Join(destPath, destFilenameInfluxDb)

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

	destFileInfluxDb, err := os.Create(destFile2InfluxDb)
	if err != nil {
		log.Fatal(err)
	}
	defer destFileInfluxDb.Close()

	csvReader := csv.NewReader(srcFile)
	csvWriter := csv.NewWriter(destFile)
	defer csvWriter.Flush()
	txtWriter := bufio.NewWriter(destFileInfluxDb)
	defer txtWriter.Flush()
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
		timestampUnix := timestampParsed.UTC().Unix()
		uuid := row[1]

		// ToDo: Das geht besser in Go
		for i := 0; i < len(c.Stations); i++ {
			station := c.Stations[i]
			if uuid == station.Id {
				parseLine(csvWriter, txtWriter, row, timestampString, timestampUnix, station)
			}
		}
	}

	// ToDo: Add Brennpaste
}

func parseLine(csvWriter *csv.Writer, txtWriter *bufio.Writer, row []string, timestampString string, timestampUnix int64, station Station) {
	diesel := row[2]
	e5 := row[3]
	e10 := row[4]
	dieselchange := row[5]
	e5change := row[6]
	e10change := row[7]

	details := []string{timestampString, station.Brand, station.City, station.Street}

	// 0=keine ??nderung, 1=??nderung, 2=Entfernt, 3=Neu
	// ToDo -1 und 2 beachten

	// Hint: Manchmal ist bei einer ??nderung (1) der Preis -0.001.
	// Keine Ahnung warum, aber die Preise werden ignoriert.

	if dieselchange == "1" && diesel[0] != '-' {
		csvWriter.Write(append(details, "Diesel", diesel))
		// ToDo Das geht besser
		txtWriter.WriteString("kraftstoffpreise,marke=" + strings.Replace(station.Brand, " ", "\\ ", -1) + ",ort=" + strings.Replace(station.City, " ", "\\ ", -1) + ",strasse=" + strings.Replace(station.Street, " ", "\\ ", -1) + ",")
		txtWriter.WriteString("sorte=Diesel preis=" + diesel + " " + strconv.FormatInt(timestampUnix, 10) + "\n")
	}
	if e5change == "1" && e5[0] != '-' {
		csvWriter.Write(append(details, "E5", e5))
		// ToDo Das geht besser
		txtWriter.WriteString("kraftstoffpreise,marke=" + strings.Replace(station.Brand, " ", "\\ ", -1) + ",ort=" + strings.Replace(station.City, " ", "\\ ", -1) + ",strasse=" + strings.Replace(station.Street, " ", "\\ ", -1) + ",")
		txtWriter.WriteString("sorte=E5 preis=" + e5 + " " + strconv.FormatInt(timestampUnix, 10) + "\n")
	}
	if e10change == "1" && e10[0] != '-' {
		csvWriter.Write(append(details, "E10", e10))
		// ToDo Das geht besser
		txtWriter.WriteString("kraftstoffpreise,marke=" + strings.Replace(station.Brand, " ", "\\ ", -1) + ",ort=" + strings.Replace(station.City, " ", "\\ ", -1) + ",strasse=" + strings.Replace(station.Street, " ", "\\ ", -1) + ",")
		txtWriter.WriteString("sorte=E10 preis=" + e10 + " " + strconv.FormatInt(timestampUnix, 10) + "\n")
	}
}
