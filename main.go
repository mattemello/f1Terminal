package main

import (
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/mattemello/f1Terminal/internal/errorsh"
	data "github.com/mattemello/f1Terminal/internal/takeData"

	tui "github.com/mattemello/f1Terminal/interface/tui"
)

func newTableRow(str [][]string) []table.Row {
	return []table.Row{
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

}

func startError() {
	file := errorsh.OpenFileLog()
	log.SetOutput(file)
	defer file.Close()
}

func main() {
	startError()

	data.TakeDriverInSession()
	on, typeSession := data.IsSessionOn()

	p := Start(typeSession)

	if on {
		Timer(p)
	} else {
		str := data.NoSession()

		tableRow := newTableRow(str)
		ms := tui.MsgUpdate{SessionOn: false, Table: tableRow}
		p.Send(ms)
	}

	select {}
}

func Start(typeSession string) *tea.Program {
	cir := data.TakeCircuit()

	cirString := tui.Circuit{
		GranprixName:    cir.CircuitShortName,
		GranprixOffName: tui.CutOff(cir.MeetingOfficialName),
		CountryName:     cir.CountryName,
		Date:            cir.DateStart,
		Location:        cir.Location,
		TypeSession:     typeSession,
	}

	p := tea.NewProgram(tui.NewModel(cirString))

	go func() {
		if _, err := p.Run(); err != nil {
			errorsh.AssertNilTer(err, "The bubbletea program run in a error")
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

			tableRow := newTableRow(str)

			ms := tui.MsgUpdate{SessionOn: true, Table: tableRow}
			p.Send(ms)
		}()
	}

}
