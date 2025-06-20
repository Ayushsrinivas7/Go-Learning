package main

import (
	"log"
     "fmt"
	"project/validator" // update to your actual module name
)

func main() {


	// Initialize the pincode validator
	validatorService, err := validator.NewCSVValidator("data/pincodes.csv")
	if err != nil {
		log.Fatalf("Failed to initialize validator: %v", err)
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
