package main

import (
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/table"
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

		tableRow := []table.Row{
			str[0],
			str[1],
			str[2],
			str[3],
			str[4],
			str[5],
			str[6],
			str[7],
			str[8],
			str[9],
			str[10],
			str[11],
			str[12],
			str[13],
			str[14],
			str[15],
			str[16],
			str[17],
			str[18],
			str[19],
		}

		p.Send(tui.MsgUpdate(tableRow))
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
			_ = data.TickedDone()
		}()
	}

}
