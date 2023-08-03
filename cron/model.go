package cron

const (
	minYear = 1970
	maxYear = 2099
)

// CronFields...
type CronFields struct {
	Seconds, Minutes, Hours, DayOfMonth, Month, DayOfWeek, Year []int
}
