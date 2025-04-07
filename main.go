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
	p := Start()

	if data.IsSessionOn() {
		Timer(p)
	} else {
		str := data.NoSession()
		p.Send(tui.MsgUpdate(str))
	}

	select {}
}

func Start() *tea.Program {
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

	return p

}

func Timer(p *tea.Program) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		go func() {
			str := data.TickedDone()
			p.Send(tui.MsgUpdate(str))
		}()
	}

}
