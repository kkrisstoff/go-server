package csv

import (
	"encoding/csv"
	"os"
)

// TODO: reuse this type
type item struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

// WriteData write data into CSV
func WriteData() {

	file, err := os.OpenFile("../../static/test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}

	x := []string{"Country", "City", "Population"}
	y := []string{"Japan", "Tokyo", "923456"}
	z := []string{"Australia", "Sydney", "789650"}
	csvWriter := csv.NewWriter(file)
	strWrite := [][]string{x, y, z}
	csvWriter.WriteAll(strWrite)
	csvWriter.Flush()
}
