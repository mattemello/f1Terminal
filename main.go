package main

import (
	"fmt"
	"time"
)

type TimeNow struct {
	Date time.Time
}

var Time TimeNow

func main() {
	Timer()
}

func Timer() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		go takeTimeNow()
	}

}

func takeTimeNow() {
	Time.Date = time.Now().UTC()
	takeData()
}

func takeData() {
	fmt.Println(Time.Date.Format("2006-01-02 15:04:05"))
}
