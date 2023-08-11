package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeSlice(t *testing.T) {
	tests := []struct {
		Start    int
		End      int
		Expected []int
	}{
		{1, 5, []int{1, 2, 3, 4, 5}},
		{0, 0, []int{0}},
		{-3, 2, []int{-3, -2, -1, 0, 1, 2}},
	}

	for _, test := range tests {
		result := RangeSlice(test.Start, test.End)
		assert.Equal(t, test.Expected, result)
	}
}

func BenchmarkRangeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RangeSlice(1, 1000)
	}
}
