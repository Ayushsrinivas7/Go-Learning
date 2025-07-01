package unit_test

func CheckDivisibilty(input int) string {

	if input%3 == 0 {
		return "THREE"
	}

	if input%5 == 0 {
		return "FIVE"
	}

	if input%2 == 0 {
		return "TWO"
	}

	return "Invalid Number"
}
