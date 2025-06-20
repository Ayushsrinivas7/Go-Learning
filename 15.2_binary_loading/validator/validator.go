package validator

import (
	"embed"
	"encoding/csv"
	"fmt"
	"strings"
)

//go:embed data/pincodes.csv
var csvFile embed.FS


type PincodeValidator interface {
	IsValid(pincode string) bool
}

type csvValidator struct {
	data map[string]bool

}

func NewEmbeddedCSVValidator() (PincodeValidator, error) {
	fileBytes, err := csvFile.ReadFile("data/pincodes.csv")

	if err != nil {
		return nil, fmt.Errorf("failed to read embedded CSV: %w", err)
	}

	reader := csv.NewReader(strings.NewReader(string(fileBytes)))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to parse CSV: %w", err)
	}

	data := make(map[string]bool)
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}
		pincode := strings.TrimSpace(record[0])
		if pincode != "" {
			data[pincode] = true
		}
	}

	return &csvValidator{data: data}, nil
}

func (v *csvValidator) IsValid(pincode string) bool {
	return v.data[pincode]
}
