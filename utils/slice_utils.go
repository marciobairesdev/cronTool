package utils

// RangeSlice...
func RangeSlice(start, end int) []int {
	slice := make([]int, 0)
	for i := start; i <= end; i++ {
		slice = append(slice, i)
	}
	return slice
}

// Contains...
func Contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
