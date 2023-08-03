package main

import (
	"fmt"
	"os"

	"github.com/marciobairesdev/cronTool/cron"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "-s" {
		fmt.Println("!!! Usage: cronTool -s \"<cron_schedule>\"")
		os.Exit(1)
	}

	fmt.Println("#################### Welcome to cronTool! ####################")
	cron.Run(os.Args[2])
}
