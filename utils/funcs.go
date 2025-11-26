package utils

// returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// returns the number of occurrences of target in slice.
func Count[T comparable](slice []T, target T) int {
	count := 0
	for _, item := range slice {
		if item == target {
			count++
		}
	}
	return count
}
