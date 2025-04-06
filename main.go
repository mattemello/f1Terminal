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

	if data.IsSessionOn() {
		Timer()
	} else {
		for {

		}
	}

}

func Start() {
	cir, err := data.TakeCircuit()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	cirString := tui.Circuit{
		GranprixName:    cir.CircuitShortName,
		GranprixOffName: cir.MeetingOfficialName,
		CountryName:     cir.CountryName,
		Date:            cir.DateStart,
		Location:        cir.Location,
	}

	p := tea.NewProgram(tui.NewModel(cirString))

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
