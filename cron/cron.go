package cron

import (
	"os"
	"os/signal"
	"time"

	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

// CronJob is a method to fire when a cron execution time is reached
type CronJob func()

// JobStatus...
type JobStatus uint8

const (
	Idle JobStatus = iota
	Running
	Finished
)

// Cron takes a cron expression, parses it, calculates the respective times
// to run, and triggers the given method when the exectuion times are reached.
//
// To create a new Cron, call [New(cronExpression, jobFunc)]
type Cron struct {
	Expression                                                  string         // The given cron expression.
	Seconds, Minutes, Hours, DayOfMonth, Month, DayOfWeek, Year []int          // Parsed portions of the given cron expression.
	Job                                                         CronJob        // Method to trigger when the execution time is reached.
	RunCount                                                    int64          // Number of times the job ran
	Status                                                      JobStatus      // Job status
	Signals                                                     chan os.Signal // Channel for interrupt signals
}

// Run triggers Cron.Job at the scheduled time
func (c *Cron) Run() {
	now := time.Now()
	nextExecutionTime := c.GetNextExecutionTime(now)
	c.Status = Running
	slog.Info("Job scheduled...", "nextExecutionTime", nextExecutionTime)
	ticker := time.NewTicker(time.Until(nextExecutionTime))
	defer ticker.Stop()
	signal.Notify(c.Signals, os.Interrupt)

	for {
		select {
		case <-ticker.C:
			c.Job()
			c.RunCount++
			now = time.Now()
			if !c.hasNextExecutionTime(now) {
				c.Status = Finished
				c.Signals <- os.Interrupt
				return
			}
			nextExecutionTime = c.GetNextExecutionTime(now)
			ticker.Reset(time.Until(nextExecutionTime))
			slog.Info("Job rescheduled...", "nextExecutionTime", nextExecutionTime)
		case <-c.Signals:
			c.Status = Finished
			return
		}
	}
}

// GetNextExecutionTime gets the next execution time for Cron based on the given time
func (c *Cron) GetNextExecutionTime(now time.Time) time.Time {
	for {
		now = now.Add(time.Second)
		if c.isNextExecutionTime(now) {
			return now
		}
	}
}

func (c *Cron) isNextExecutionTime(now time.Time) bool {
	return slices.Contains(c.Seconds, now.Second()) &&
		slices.Contains(c.Minutes, now.Minute()) &&
		slices.Contains(c.Hours, now.Hour()) &&
		slices.Contains(c.DayOfMonth, now.Day()) &&
		slices.Contains(c.Month, int(now.Month())) &&
		slices.Contains(c.DayOfWeek, int(now.Weekday())) &&
		slices.Contains(c.Year, now.Year())
}

func (c *Cron) hasNextExecutionTime(now time.Time) bool {
	for _, second := range c.Seconds {
		for _, minute := range c.Minutes {
			for _, hour := range c.Hours {
				for _, day := range c.DayOfMonth {
					for _, month := range c.Month {
						for _, year := range c.Year {
							nextTime := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
							if slices.Contains(c.DayOfWeek, int(nextTime.Weekday())) && nextTime.After(now) {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}

// New creates a new Cron with the given cron expression and JobFunc.
func New(cronExpression string, job CronJob) (*Cron, error) {
	c, err := parseCronExpression(cronExpression)
	if err != nil {
		return nil, err
	}

	c.Job = job
	c.RunCount = 0
	c.Status = Idle
	c.Signals = make(chan os.Signal, 1)

	return c, nil
}
