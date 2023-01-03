package exercise001

func isDevisableBySeven(number int) bool {
	return number%7 == 0
}

func isNotMultipleOfFive(number int) bool {
	lastDigit := number % 10
	if lastDigit == 0 || lastDigit == 5 {
		return false
	}
	return true
}
