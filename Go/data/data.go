package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// OpenCSV opens a CSV file
func OpenCSV(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	return file, nil
}

// ReadCSV reads all records from a CSV file, skipping the header
func ReadCSV(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)

	// Read and discard the first record (header)
	_, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header: %w", err)
	}

	// Read all the remaining records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading records: %w", err)
	}

	return records, nil
}

// ConvertToFloat64 converts a 2D slice of strings to a 2D slice of float64s
func ConvertToFloat64(records [][]string) ([][]float64, error) {
	data := make([][]float64, len(records))
	for i, record := range records {
		// Ignore the last column
		data[i] = make([]float64, len(record)-1)
		for j := 0; j < len(record)-1; j++ {
			value, err := strconv.ParseFloat(record[j], 64)
			if err != nil {
				return nil, fmt.Errorf("error converting value to float64: %w", err)
			}
			data[i][j] = value
		}
	}
	return data, nil
}

// ReadData reads data from a CSV file and converts it to a 2D slice of float64s
func ReadData() ([][]float64, error) {
	file, err := OpenCSV("../iris.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := ReadCSV(file)
	if err != nil {
		return nil, err
	}

	data, err := ConvertToFloat64(records)
	if err != nil {
		return nil, err
	}

	return data, nil
}
