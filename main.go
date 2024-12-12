package main

import (		

	"github.com/jasonlvhit/gocron"
)

func task() {
	// Do your task here	
}
func main() {
	gocron.Every(15).Second().Do(task)
}
