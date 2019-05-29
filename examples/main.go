package main

import (
	simlog "github.com/OloloevReal/go-simple-log"
)

func main() {
	slog := simlog.NewLogger()
	slog.Printf("test log message")

	simlog.SetOptions(simlog.SetDebug)
	simlog.Printf("This is a [%s] log", "global")
}
