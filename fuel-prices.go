package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/4s1/fuel-prices/conf"
	"gitlab.com/4s1/fuel-prices/directory"
	"gitlab.com/4s1/fuel-prices/git"
)

func main() {
	fmt.Println("--- Fuel Prices ---")
	configFileName := "config.json"

	// Program parameters
	useInit := flag.Bool("init", false, "Initialize configuration file")
	flag.Parse()

	// Handle parameters
	if *useInit {
		err := conf.CreateNewConfigurationFile(configFileName)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Load configuration
	config, err := conf.LoadConfigurationFile(configFileName)
	if err != nil {
		panic(err)
	}

	// Clone or update repo
	// ToDo: Handle other err
	if _, err := os.Stat(config.TankerkoenigDataFolder); os.IsNotExist(err) {
		repoPath := "https://tankerkoenig@dev.azure.com/tankerkoenig/tankerkoenig-data/_git/tankerkoenig-data"
		git.Clone(config.TankerkoenigDataFolder, repoPath)
	} else {
		git.Pull(config.TankerkoenigDataFolder)
	}

	// Get input data date range
	firstSrcFile, lastSrcFile, err := directory.FirstAndLastDate(config.TankerkoenigDataFolder + "/" + "prices")
	if err != nil {
		panic(err)
	}
	fmt.Println(firstSrcFile)
	fmt.Println(lastSrcFile)

	// Get parsed data date range
	firstDestFile, lastDestFile, err := directory.FirstAndLastDate(config.CsvDataFolder)
	if err != nil {
		panic(err)
	}
	fmt.Println(firstDestFile)
	fmt.Println(lastDestFile)
}
