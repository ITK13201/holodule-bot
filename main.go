package main

import (
	"github.com/ITK13201/holodule-bot/processes"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("Usage: require 1 argument ('daily' or 'coming-soon'])")
	}

	switch args[1] {
	case "daily":
		processes.NotifyDaily()
	case "coming-soon":
		processes.NotifyComingSoon()
	default:
		panic("Usage: require 1 argument ('daily' or 'coming-soon'])")
	}
}
