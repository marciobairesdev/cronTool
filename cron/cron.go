package cron

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// Run...
func Run(cronExpression string) {
	// Create a ticker that fires every second.
	ticker := time.NewTicker(time.Second)
	nextExecutionDate := getNextExecutionDate(cronExpression)

	// Create a buffered channel of type chan os.Signal.
	c := make(chan os.Signal, 1)

	// Listen for signals.
	signal.Notify(c, os.Interrupt)

	// Loop forever, firing the action when the cron schedule is met.
	for {
		select {
		case <-ticker.C:
			if nextExecutionDate.After(time.Now()) {
				continue
			}

			// Trigger the action.
			fmt.Println("\t# cronTool schedule triggered!")
			nextExecutionDate = getNextExecutionDate(cronExpression)
		case signal := <-c:
			// We've received a signal, so exit.
			fmt.Println("\n\n!!! cronTool is exiting due to signal:", signal)
			os.Exit(0)
		}
	}
}

func getNextExecutionDate(cronExpression string) (nextExecutionDate time.Time) {
	nextExecutionDate, err := GetNextDate(cronExpression)
	if err != nil {
		fmt.Println("\t!!! Error:", err)
		os.Exit(1)
	}
	fmt.Println("\n# Next execution date:", nextExecutionDate.Format(time.DateTime))
	return
}
