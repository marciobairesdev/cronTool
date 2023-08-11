package cron

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCronFieldMinMaxValues(t *testing.T) {
	tests := []struct {
		FieldType CronFieldType
		MinValue  int
		MaxValue  int
	}{
		{Second, 0, 59},
		{Minute, 0, 59},
		{Hour, 0, 23},
		{DayOfMonth, 1, 31},
		{Month, int(time.January), int(time.December)},
		{DayOfWeek, int(time.Sunday), int(time.Saturday)},
		{Year, 1970, 2099},
	}

	for _, test := range tests {
		min, max := getCronFieldMinMaxValues(test.FieldType)
		assert.Equal(t, test.MinValue, min)
		assert.Equal(t, test.MaxValue, max)
	}
}
