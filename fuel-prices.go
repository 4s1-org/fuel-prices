package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("--- Fuel Prices ---")
	configFileName := "config.json"

	// Program parameters
	useInit := flag.Bool("init", false, "Initialize configuration file")
	flag.Parse()

	// Handle parameters
	if *useInit {
		err := CreateNewConfigurationFile(configFileName)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Load configuration
	config, err := LoadConfigurationFile(configFileName)
	if err != nil {
		panic(err)
	}

	// Clone or update repo
	// ToDo: Handle other err
	if _, err := os.Stat(config.TankerkoenigDataFolder); os.IsNotExist(err) {
		repoPath := "https://tankerkoenig@dev.azure.com/tankerkoenig/tankerkoenig-data/_git/tankerkoenig-data"
		Clone(config.TankerkoenigDataFolder, repoPath)
	} else {
		Pull(config.TankerkoenigDataFolder)
	}

	// Get input data date range
	firstSrcFile, lastSrcFile, err := FirstAndLastDate(config.TankerkoenigDataFolder + "/" + "prices")
	if err != nil {
		panic(err)
	}
	fmt.Println("firstSrcFile:", firstSrcFile)
	fmt.Println("firstSrcFile:", lastSrcFile)

	// Get parsed data date range
	firstDestFile, lastDestFile, err := FirstAndLastDate(config.CsvDataFolder)
	if err != nil {
		panic(err)
	}
	fmt.Println("firstDestFile:", firstDestFile)
	fmt.Println("lastDestFile:", lastDestFile)

	start := lastDestFile
	if firstDestFile.Before(firstSrcFile) {
		start = firstSrcFile
	}
	end := lastSrcFile

	fmt.Println("start:", start)
	fmt.Println("end:", end)

	// Parse
	Parse(start, end, config)
}
