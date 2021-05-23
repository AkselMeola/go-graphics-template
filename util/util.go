package util

// Limita given value to given limits
func Limit(value int, lowerLimit, upperLimit int) int {
	if value > upperLimit {
		return upperLimit
	}
	if value < lowerLimit {
		return lowerLimit
	}

	return value
}
