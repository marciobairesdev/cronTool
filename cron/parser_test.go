package cron

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCronExpression_ValidExpression(t *testing.T) {
	tests := []struct {
		Expression string
	}{
		{"* * * * * * *"},
		{"45 23 * * 6 * *"},
		{"   0,10,20,30,40,50 * * * * * *"},
		{"* * * * * * 2025   "},
		{"*/3 1-5 19/21 * * * 2024"},
		{"* *      * * 10    4-6 *"},
		{"57 0 10 8 * * *"},
		{"8 * * * * * 1970"},
		{"0 * * * * * 2099"},
		{"0 * * * * * 2099"},
	}

	for _, test := range tests {
		c, err := parseCronExpression(test.Expression)
		assert.NotNil(t, c)
		assert.Nil(t, err)
	}
}

func TestParseCronExpression_InvalidExpression(t *testing.T) {
	tests := []struct {
		Expression string
	}{
		{""},
		{" "},
		{"        "},
		{"a invalid * * * * 2023"},
		{"* * *"},
		{"* * * * *"},
		{"* * * * * *"},
		{"* * * * * * * * * * *"},
		{"*-1 * * * * * *"},
		{"? * * * * *    2025"},
		{"2-3/10 * * * * * * *   "},
		{"-1 * * * * * *"},
		{"60 * * * * * *"},
		{"* 60 * * * * *"},
		{"* * 24 * * * *"},
		{"* * * 32 * * *"},
		{"* * * * 13 * *"},
		{"* * * * * 7 *"},
		{"* * * * * * 1969"},
		{"* * * * * * 3000"},
	}

	for _, test := range tests {
		c, err := parseCronExpression(test.Expression)
		assert.Nil(t, c)
		assert.NotNil(t, err)
	}
}

func TestSanitizeSpaces(t *testing.T) {
	tests := []struct {
		Input          string
		ExpectedOutput string
	}{
		{"    *   *    *     * * * *", "* * * * * * *"},
		{"* * * * *    * *    ", "* * * * * * *"},
		{" * * * * * * * ", "* * * * * * *"},
	}

	for _, test := range tests {
		output := sanitizeSpaces(test.Input)
		assert.Equal(t, test.ExpectedOutput, output)
	}
}

func TestParseCronField(t *testing.T) {
	tests := []struct {
		Field          string
		ExpectedOutput []int
		FieldType      CronFieldType
	}{
		{"*/20", []int{0, 20, 40}, Second},
		{"1,5,29", []int{1, 5, 29}, Minute},
		{"2/3", []int{2, 5, 8, 11, 14, 17, 20, 23}, Hour},
		{"31", []int{31}, DayOfMonth},
		{"2-6", []int{2, 3, 4, 5, 6}, Month},
		{"0-6", []int{0, 1, 2, 3, 4, 5, 6}, DayOfWeek},
		{"2023-2025", []int{2023, 2024, 2025}, Year},
	}

	for _, test := range tests {
		output := parseCronField(test.Field, test.FieldType)
		assert.Equal(t, test.ExpectedOutput, output)
	}
}

func BenchmarkParseCronField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseCronField("1970-2030", Year)
	}
}

func BenchmarkSanitizeSpaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sanitizeSpaces("    *   *    *     * * * *")
	}
}
