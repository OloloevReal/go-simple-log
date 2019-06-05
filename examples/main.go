package main

import (
	"bytes"
	"fmt"

	simlog "github.com/OloloevReal/go-simple-log"
)

func main() {

	def := simlog.GetDafault()
	var buf bytes.Buffer
	def.SetOutput(&buf)

	simlog.Printf("This is a INFO [%s] log", "global")
	simlog.SetOptions(simlog.SetDebug)
	//simlog.SetOptions(simlog.SetDebug, simlog.SetCaller)
	simlog.Printf("[DEBUG] This is a [%s] log", "global")
	simlog.Println("test Println")
	//simlog.Fatalln("[FATAL] this is Fatalln")
	//simlog.Panicln("[PANIC] this is Panicln")
	simlog.Println("test Println1")
	fmt.Println(buf.String())
	simlog.Fatalln("[FATAL] this is Fatalln")

	//simlog.Panicf("[INFO] This is a panic log")

}
