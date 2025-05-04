package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func validateCSV(filename string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return false, fmt.Errorf("error reading CSV: %w", err)
	}

	if len(records) <= 1 {
		return false, fmt.Errorf("file is empty or only has header")
	}

	header := records[0]
	numColumns := len(header)
	
	for i, row := range records {
        if i == 0{
            continue
        }
		if len(row) > numColumns {
			return false, fmt.Errorf("row %d has more columns than header", i+1)
		}
		for j, headerColumn := range header {
			if j >= len(row) {
				return false, fmt.Errorf("row %d is missing column %s", i+1, headerColumn)
			}
			if strings.TrimSpace(row[j]) == ""{
                return false, fmt.Errorf("row %d in column %s is empty", i+1, headerColumn)
            }
		}
	}

	return true, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	isValid, err := validateCSV(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if isValid {
		fmt.Println("CSV is valid")
	} else {
		fmt.Println("CSV is not valid")
	}
}