package conf

import (
	"encoding/json"
	"errors"
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
	Stations               []Station `json:"stations"`
	TankerkoenigDataFolder string    `json:"tankerkoenigDataFolder"`
	CsvDataFolder          string    `json:"csvDataFolder"`
}

func LoadConfigurationFile(configFileName string) (*Configuration, error) {
	_, err := os.Stat(configFileName)
	if errors.Is(err, os.ErrNotExist) {
		// Config doesn't exists
		return nil, errors.New("No configuration file found")
	}
	if err != nil {
		panic(err)
	}

	file, err := ioutil.ReadFile(configFileName)
	if err != nil {
		panic(err)
	}
	configuration := Configuration{}
	err = json.Unmarshal([]byte(file), &configuration)
	if err != nil {
		panic(err)
	}

	return &configuration, nil
}

// Erzeut eine neue Konfigurationdatei mit Dummy-Daten.
func CreateNewConfigurationFile(configFileName string) error {
	_, err := os.Stat(configFileName)
	if err == nil {
		// Config exists
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
