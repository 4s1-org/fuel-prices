package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func FirstAndLastDate(folder string) (time.Time, time.Time, error) {
	_, err := ioutil.ReadDir(folder)
	if err != nil {
		return time.Now(), time.Now(), err
	}

	// Year
	minYear, maxYear, err := minAndMaxFoldername(folder)
	if err != nil {
		log.Fatal(err)
	}
	if minYear == 9999 || maxYear == 0 {
		return time.Date(1970, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
			time.Date(2035, time.Month(12), 31, 0, 0, 0, 0, time.UTC),
			nil
	}

	// Month
	minMonth, _, err := minAndMaxFoldername(folder + "/" + fmt.Sprintf("%04d", minYear))
	_, maxMonth, err := minAndMaxFoldername(folder + "/" + fmt.Sprintf("%04d", maxYear))

	// Day
	minDay, _, err := minAndMaxFilename(folder + "/" + fmt.Sprintf("%04d", minYear) + "/" + fmt.Sprintf("%02d", minMonth))
	_, maxDay, err := minAndMaxFilename(folder + "/" + fmt.Sprintf("%04d", maxYear) + "/" + fmt.Sprintf("%02d", maxMonth))

	// Return
	return time.Date(minYear, time.Month(minMonth), minDay, 0, 0, 0, 0, time.UTC),
		time.Date(maxYear, time.Month(maxMonth), maxDay, 0, 0, 0, 0, time.UTC),
		nil
}

func minAndMaxFoldername(folder string) (int, int, error) {
	items, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	if len(items) == 0 {
		return 0, 0, errors.New("No items in folder: " + folder)
	}

	minValue := 9999
	maxValue := 0

	for _, item := range items {
		if !item.IsDir() {
			continue
		}

		foldername := item.Name()
		if foldername == ".git" {
			continue
		}

		i, err := strconv.Atoi(foldername)
		if err != nil {
			log.Fatal(err)
		}

		if i < minValue {
			minValue = i
		}
		if i > maxValue {
			maxValue = i
		}
	}

	return minValue, maxValue, nil
}

func minAndMaxFilename(folder string) (int, int, error) {
	items, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	if len(items) == 0 {
		return 0, 0, errors.New("No items in folder: " + folder)
	}

	minValue := 9999
	maxValue := 0

	for _, item := range items {
		if item.IsDir() {
			continue
		}

		filename := item.Name()

		i, err := strconv.Atoi(filename[8:10])
		if err != nil {
			log.Fatal(err)
		}

		if i < minValue {
			minValue = i
		}
		if i > maxValue {
			maxValue = i
		}
	}

	return minValue, maxValue, nil
}
