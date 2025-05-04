package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	folderName := "csv"
	fileName := "requi_ot_global.csv"
	headers := []string{"requi_id", "requi_code", "requi_cat", "requi_text"}

	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	filePath := filepath.Join(folderName, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(headers)
	if err != nil {
		fmt.Println("Error writing headers:", err)
		return
	}

	fmt.Println("CSV file created successfully at:", filePath)
}
