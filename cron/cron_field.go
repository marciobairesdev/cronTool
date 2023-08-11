package cron

import "time"

// CronFieldType...
type CronFieldType uint8

const (
	Second CronFieldType = iota
	Minute
	Hour
	DayOfMonth
	Month
	DayOfWeek
	Year
)

func getCronFieldMinMaxValues(fieldType CronFieldType) (int, int) {
	switch fieldType {
	case Second, Minute:
		return 0, 59
	case Hour:
		return 0, 23
	case DayOfMonth:
		return 1, 31
	case Month:
		return int(time.January), int(time.December)
	case DayOfWeek:
		return int(time.Sunday), int(time.Saturday)
	default:
		return 1970, 2099
	}
}
