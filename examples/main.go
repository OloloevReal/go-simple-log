package main

import (
	simlog "github.com/OloloevReal/go-simple-log"
)

func main() {
	slog := simlog.NewLogger()
	slog.Printf("[ERROR] test log message")

	simlog.Printf("This is a INFO [%s] log", "global")
	//simlog.SetOptions(simlog.SetDebug)
	simlog.SetOptions(simlog.SetDebug, simlog.SetCaller)
	simlog.Printf("[DEBUG] This is a [%s] log", "global")

	simlog.Fatalf("[FATAL] This is a fatal log")
	//simlog.Panicf("[INFO] This is a panic log")
}
