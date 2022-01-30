package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Station struct {
	Id     string `json:"id"`
	Brand  string `json:"brand"`
	City   string `json:"city"`
	Street string `json:"street"`
}

type Configuration struct {
	Stations []Station `json:"stations"`
}

func main() {
	fmt.Println("--- Fuel Prices ---")

	useInit := flag.Bool("init", false, "Initalize configuration file")
	flag.Parse()

	if *useInit {
		err := createNewConfigurationFile()
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Erzeut eine neue Konfigurationdatei mit Dummy-Daten.
func createNewConfigurationFile() error {
	configFileName := "config.json"

	_, err := os.Stat(configFileName)
	if err == nil {
		// Datei existiert
		return errors.New("Configuration already exists")
	}
	if errors.Is(err, os.ErrNotExist) {
		config := createExampleData()
		file, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(configFileName, file, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}

	return nil
}

// Erstelle einen Configuration-Struct mit Dummy-Werten.
func createExampleData() Configuration {
	station := Station{
		Id:     "uuid",
		Brand:  "brand",
		City:   "city",
		Street: "street",
	}
	config := Configuration{
		Stations: []Station{station},
	}

	return config
}
