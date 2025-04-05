package main

import (
	"time"

	data "github.com/mattemello/f1Terminal/internal/takeData"
)

func main() {
	Timer()
}

func Timer() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		go data.StartTicked()
	}

}
