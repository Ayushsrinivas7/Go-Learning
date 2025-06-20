package validator

import (
	"encoding/csv"
	"fmt"
	"os"
)

type PincodeValidator interface {
	IsValid(pincode string) bool
}

type csvValidator struct {
	data map[string]bool
}

func (v *csvValidator) IsValid(pincode string) bool {
	return v.data[pincode]
}



// NewCSVValidator creates a new validator from the given CSV file
func NewCSVValidator(filePath string) (PincodeValidator, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	data := make(map[string]bool)
	for i, record := range records {
		if i == 0 {
			continue // skip header
		}
		data[record[0]] = true // record[0] is the pincode
	}

	return &csvValidator{data: data}, nil
}
