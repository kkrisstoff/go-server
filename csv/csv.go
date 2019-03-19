package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// TODO: reuse this type
type item struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

// WriteData write data into CSV
func WriteData(name string, message string) {
	fmt.Printf("Name: %s\nMessage: %s\n", name, message)

	file, err := os.OpenFile("static/test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	data := []string{"123", name, message}

	csvWriter := csv.NewWriter(file)
	strWrite := [][]string{data}

	csvWriter.WriteAll(strWrite)
	csvWriter.Flush()
}

// ReadData read data into CSV
func ReadData() {
	file, err := os.Open("static/test.csv")
	if err != nil {
		fmt.Printf("Couldn't open the csv file,  %s/n", err.Error())
	}
	defer file.Close()
	// Parse the file
	reader := csv.NewReader(file)
	// NOTE: '\t': tab separator, ',': by default
	// reader.Comma = '\t'
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading csv file :  %s", err.Error())
			break
		}

		fmt.Println(record)
	}

}
