package cron

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/marciobairesdev/cronTool/utils"
)

const cronRegExp = `^($|\s*=|(\*|(?:\*)(?:(?:|\/)(?:[0-5]?\d))|(?:[0-5]?\d)(?:(?:-|\/|\,)(?:[0-5]?\d))?(?:,(?:[0-5]?\d)(?:(?:-|\/|\,)(?:[0-5]?\d))?)*)\s+(\*|(?:\*)(?:(?:|\/)(?:[0-5]?\d))|(?:[0-5]?\d)(?:(?:-|\/|\,)(?:[0-5]?\d))?(?:,(?:[0-5]?\d)(?:(?:-|\/|\,)(?:[0-5]?\d))?)*)\s+(\*|(?:\*)(?:(?:|\/)(?:[01]?\d|2[0-3]))|(?:[01]?\d|2[0-3])(?:(?:-|\/|\,)(?:[01]?\d|2[0-3]))?(?:,(?:[01]?\d|2[0-3])(?:(?:-|\/|\,)(?:[01]?\d|2[0-3]))?)*)\s+(\*|(?:\*)(?:(?:\/)(?:0?[1-9]|[12]\d|3[01]))|(?:0?[1-9]|[12]\d|3[01])(?:(?:-|\/|\,)(?:0?[1-9]|[12]\d|3[01]))?(?:,(?:0?[1-9]|[12]\d|3[01])(?:(?:-|\/|\,)(?:0?[1-9]|[12]\d|3[01]))?)*)\s+(\*|(?:\*)(?:(?:\/)(?:[1-9]|1[012]))|(?:[1-9]|1[012])(?:(?:-|\/|\,)(?:[1-9]|1[012]))?(?:,(?:[1-9]|1[012])(?:(?:-|\/|\,)(?:[1-9]|1[012]))?)*|\*)\s+(\*|(?:\*)(?:(?:\/)(?:[0-6]))|(?:[0-6])(?:(?:-|\/|\,)(?:[0-6]))?(?:,(?:[0-6])(?:(?:-|\/|\,)(?:[0-6]))?)*|\**)\s+(\*|(?:\*)(?:(?:\/)(?:19[7-9]\d|20\d{2}))|(?:19[7-9]\d|20\d{2})(?:(?:-|\/|\,)(?:19[7-9]\d|20\d{2}))?(?:,(?:19[7-9]\d|20\d{2})(?:(?:-|\/|\,)(?:19[7-9]\d|20\d{2}))?)*))$`

func parseCronExpression(cronExpression string) (*Cron, error) {
	expr := sanitizeSpaces(cronExpression)
	if expr == "" || !regexp.MustCompile(cronRegExp).MatchString(expr) {
		return nil, fmt.Errorf("invalid cron expression")
	}

	fields := strings.Fields(expr)
	return &Cron{
		Expression: expr,
		Seconds:    parseCronField(fields[0], Second),
		Minutes:    parseCronField(fields[1], Minute),
		Hours:      parseCronField(fields[2], Hour),
		DayOfMonth: parseCronField(fields[3], DayOfMonth),
		Month:      parseCronField(fields[4], Month),
		DayOfWeek:  parseCronField(fields[5], DayOfWeek),
		Year:       parseCronField(fields[6], Year),
	}, nil
}

func sanitizeSpaces(s string) string {
	return strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(s, " "))
}

func parseCronField(field string, fieldType CronFieldType) (values []int) {
	minVal, maxVal := getCronFieldMinMaxValues(fieldType)

	if field == "*" {
		return utils.RangeSlice(minVal, maxVal)
	}

	parts := strings.Split(field, ",")
	for _, part := range parts {
		switch {
		case strings.Contains(part, "-"):
			rangeParts := strings.Split(part, "-")
			start, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			values = append(values, utils.RangeSlice(start, end)...)
		case strings.Contains(part, "/"):
			incrementParts := strings.Split(part, "/")
			start, _ := strconv.Atoi(incrementParts[0])
			increment, _ := strconv.Atoi(incrementParts[1])
			for i := start; i <= maxVal; i += increment {
				values = append(values, i)
			}
		default:
			val, _ := strconv.Atoi(part)
			values = append(values, val)
		}
	}
	return
}
