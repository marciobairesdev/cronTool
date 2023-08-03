package cron

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/marciobairesdev/cronTool/utils"
)

// ParseCronExpression...
func ParseCronExpression(expression string) (CronFields, error) {
	expr := strings.TrimSpace(expression)
	if expr == "" {
		return CronFields{}, fmt.Errorf("empty cron expression")
	}

	fields := strings.Fields(expr)
	numParts := len(fields)
	if numParts < 5 || numParts > 7 {
		return CronFields{}, fmt.Errorf("invalid cron expression")
	}

	parsed := CronFields{}

	yearIndex := -1
	if numParts == 6 && regexp.MustCompile(`\d{4}$`).MatchString(fields[5]) {
		yearIndex = 5
		numParts = 7
	}

	parsed.Seconds = parseField(fields[0], 0, 59)
	parsed.Minutes = parseField(fields[1], 0, 59)
	parsed.Hours = parseField(fields[2], 0, 23)
	parsed.DayOfMonth = parseField(fields[3], 1, 31)
	parsed.Month = parseField(fields[4], 1, 12)
	parsed.DayOfWeek = parseField(fields[5], 0, 6)
	if yearIndex != -1 {
		parsed.Year = parseField(fields[yearIndex], minYear, maxYear)
	} else {
		parsed.Year = parseField("*", minYear, maxYear)
	}

	return parsed, nil
}

func parseField(field string, minVal, maxVal int) []int {
	if field == "*" {
		return utils.RangeSlice(minVal, maxVal)
	}

	values := make([]int, 0)
	parts := strings.Split(field, ",")
	for _, part := range parts {
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			start, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			values = append(values, utils.RangeSlice(start, end)...)
		} else if strings.Contains(part, "/") {
			incrementParts := strings.Split(part, "/")
			start, _ := strconv.Atoi(incrementParts[0])
			increment, _ := strconv.Atoi(incrementParts[1])
			for i := start; i <= maxVal; i += increment {
				values = append(values, i)
			}
		} else {
			val, _ := strconv.Atoi(part)
			values = append(values, val)
		}
	}

	// Filter out values outside the allowed range
	filtered := make([]int, 0)
	for _, v := range values {
		if v >= minVal && v <= maxVal {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
