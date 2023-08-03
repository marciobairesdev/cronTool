package cron

import (
	"fmt"
	"time"

	"github.com/marciobairesdev/cronTool/utils"
)

// GetNextDate...
func GetNextDate(cronExpression string) (time.Time, error) {
	cron, err := ParseCronExpression(cronExpression)
	if err != nil {
		return time.Time{}, err
	}

	now := time.Now()
	currentYear := now.Year()

	for currentYear >= minYear && currentYear <= maxYear {
		now = now.Add(time.Second)
		if !utils.Contains(cron.Seconds, now.Second()) {
			continue
		}
		if !utils.Contains(cron.Minutes, now.Minute()) {
			continue
		}
		if !utils.Contains(cron.Hours, now.Hour()) {
			continue
		}
		if !utils.Contains(cron.DayOfMonth, now.Day()) {
			continue
		}
		if !utils.Contains(cron.Month, int(now.Month())) {
			continue
		}
		if !utils.Contains(cron.DayOfWeek, int(now.Weekday())) {
			continue
		}
		if len(cron.Year) > 0 && !utils.Contains(cron.Year, now.Year()) {
			continue
		}

		return now, nil
	}

	return time.Time{}, fmt.Errorf("no valid execution date found for the given cron expression")
}
