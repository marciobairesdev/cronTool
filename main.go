package main

import (
	"os"
	"strings"

	"github.com/marciobairesdev/cronTool/cron"
	"golang.org/x/exp/slog"
)

func main() {
	slog.Info("Welcome to cronTool!")

	if len(os.Args) != 3 || os.Args[1] != "-s" || strings.TrimSpace(os.Args[2]) == "" {
		slog.Warn("Usage: cronTool -s '<cron_schedule>'.")
		os.Exit(1)
	}

	c, err := cron.New(os.Args[2], func() {
		slog.Info("Job triggered!")
	})
	if err != nil {
		slog.Error("Cron error...", "details", err)
		os.Exit(1)
	}

	go c.Run()

	<-c.Signals
	slog.Info("cronTool is exiting...")
}
