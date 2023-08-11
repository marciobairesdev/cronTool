package utils

// RangeSlice generates a slice with from start to end
func RangeSlice(start, end int) (slice []int) {
	for i := start; i <= end; i++ {
		slice = append(slice, i)
	}
	return
}
