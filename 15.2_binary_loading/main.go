package main

import (
	"log"
	"fmt"
	"project/validator" // Replace "your_project" with your actual module name
)

func main() {
	

	// Initialize embedded pincode validator
	validatorService, err := validator.NewEmbeddedCSVValidator()
	if err != nil {
		log.Fatalf("Failed to load pincodes: %v", err)
	}

	testPostcodes := []string{"110000", "560001", "999999"}

	for _, code := range testPostcodes {
		if validatorService.IsValid(code) {
			fmt.Printf("Postcode %s is VALID ✅\n", code)
		} else {
			fmt.Printf("Postcode %s is INVALID ❌\n", code)
		}
	}

}
