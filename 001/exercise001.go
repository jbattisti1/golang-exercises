package exercise001

import (
	"strconv"
	"strings"
)

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

func findNumberDevisableBySevenButNotByFive(low, high int) string {
	var result []string
	for i := low; i <= high; i++ {
		if isDevisableBySeven(i) && isNotMultipleOfFive(i) {
			result = append(result, strconv.Itoa(i))
		}
	}
	return strings.Join(result, ",")
}
