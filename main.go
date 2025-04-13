package main

import (
	"os"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

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
	errorsh.OpenFileLog()
}

func controlSession(tyS string) tui.Circuit {
	cir := data.TakeCircuit()
	var cirString tui.Circuit
	cirString = tui.Circuit{
		GranprixName:    cir.CircuitShortName,
		GranprixOffName: tui.CutOff(cir.MeetingOfficialName),
		// GranprixOffName: tui.CutOff("FORMULA 1 MSC CRUISES GRAN PREMIO DEL MADE IN ITALY E DELL'EMILIA-ROMAGNA 2024"),
		CountryName: cir.CountryName,
		Date:        cir.DateStart,
		Location:    cir.Location,
		TypeSession: tyS,
	}

	return cirString
}

func main() {
	startError()

	data.TakeDriverInSession()
	on, typeSession := data.IsSessionOn()

	if on {
		typeSession += lipgloss.NewStyle().Foreground(lipgloss.Color("#d20f39")).Render(" *")
	}

	p := Start(typeSession)

	if on {
		Timer(p)
	} else {
		str := data.NoSession()

		tableRow := newTableRow(str)
		ms := tui.MsgUpdateTable{SessionOn: false, Table: tableRow}
		p.Send(ms)
	}

	for {
		on, typeSession = data.IsSessionOn()
		if on {
			typeSession += lipgloss.NewStyle().Foreground(lipgloss.Color("#d20f39")).Render(" *")
			p.Send(tui.MsgUpdateCiruit(controlSession(typeSession)))
			Timer(p)
		}

		time.Sleep(20 * time.Second)
	}

}

func Start(typeSession string) *tea.Program {
	p := tea.NewProgram(tui.NewModel(controlSession(typeSession)))

	go func() {
		_, err := p.Run()
		errorsh.AssertNilTer(err, "The bubbletea program run in a error")

		os.Exit(0)
	}()

	return p

}

func Timer(p *tea.Program) {
	laps := make(map[int]data.Laps)
	ticker := time.NewTicker(25 * time.Second)
	defer ticker.Stop()

	for {
		go func() {
			str := data.TickedDone(&laps)

			tableRow := newTableRow(str)

			ms := tui.MsgUpdateTable{SessionOn: true, Table: tableRow}
			p.Send(ms)
		}()
		<-ticker.C
	}
}
