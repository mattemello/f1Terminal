package main

import (
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	data "github.com/mattemello/f1Terminal/internal/takeData"
	tui "github.com/mattemello/f1Terminal/internal/tui"
)

func main() {
	Start()

	if data.isSessionOn() {
		Timer()
	}
}

func Start() {
	p := tea.NewProgram(tui.NewModel())

	go func() {
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

}

func Timer() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		go data.TickedDone()
	}

}
