package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Global map to hold postcodes
var postcodeMap map[string]bool

// Load the CSV file and fill the postcodeMap
func loadPostcodes(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read csv: %v", err)
	}

	postcodeMap = make(map[string]bool)
    count :=0 
	for i, record := range records {
		if i == 0 {
			continue // skip header
		}
		postcode := record[0]
		count +=1 
		postcodeMap[postcode] = true
	}
     println(count)
	return nil
}

// Check if the postcode exists
func IsPostcodeValid(code string) bool {
	return postcodeMap[code]
}

func main() {
	// Adjust the path if needed
	err := loadPostcodes("/Users/venkatasai/Desktop/go_learning/basics/15_loading_csv/pincodes.csv")
	if err != nil {
		fmt.Println("Failed to load pincodes:", err)
		return
	}

	// Test cases
	testPostcodes := []string{"110000", "560001", "999999"}

	for _, code := range testPostcodes {
		if IsPostcodeValid(code) {
			fmt.Printf("Postcode %s is VALID ✅\n", code)
		} else {
			fmt.Printf("Postcode %s is INVALID ❌\n", code)
		}
	}
}
