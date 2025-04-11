package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/mattemello/f1Terminal/internal/errorsh"
)

func NewModel(cir Circuit) mainModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	m := mainModel{top: cir, spinner: s}
	return m
}

func (m mainModel) Init() tea.Cmd {
	var err error
	m.Terminal.widthT, m.Terminal.heightT, err = term.GetSize(0)
	errorsh.AssertNilTer(err, "The program can't take the dimension of your terminal")

	errorsh.AssertNotAppening((m.Terminal.widthT < width || m.Terminal.heightT < heightTotal), "Error, The terminal is too small")

	return tea.Batch(tea.EnterAltScreen, m.spinner.Tick)
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case MsgUpdateTable:
		var stTable = table.DefaultStyles()

		stTable.Header = stTable.Header.BorderStyle(lipgloss.NormalBorder()).BorderBottom(true).Bold(false).Foreground(lipgloss.Color("#ed8796"))
		stTable.Selected = stTable.Selected.Bold(false).Background(lipgloss.Color("#8caaee")).Foreground(lipgloss.Color("#181926"))

		if !msg.SessionOn {
			m.Table = table.New(table.WithColumns(columsPos), table.WithRows(msg.Table), table.WithFocused(false), table.WithHeight(21))

			m.Table.SetStyles(stTable)
		} else {
			m.Table = table.New(table.WithColumns(columsSess), table.WithRows(msg.Table), table.WithFocused(false), table.WithHeight(21))

			m.Table.SetStyles(stTable)

		}
		m.tableOn = true
		break

	case MsgUpdateCiruit:
		m.top = Circuit(msg)
		break

	case MsgError:

		break

	case tea.WindowSizeMsg:

		var err error
		m.Terminal.widthT, m.Terminal.heightT, err = term.GetSize(0)
		errorsh.AssertNilTer(err, "The program can't take the dimension of your terminal")

		if m.Terminal.widthT < width || m.Terminal.heightT < heightTotal {
			m.ErrorMessage = "!!! THE TERMINAL IS TOO SMALL HELP!!!"
		} else {
			m.ErrorMessage = ""
		}

		break

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	default:
		if !m.tableOn {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}
	}

	m.Table.Focus()
	m.Table, cmd = m.Table.Update(msg)

	return m, cmd
}

func (m mainModel) View() string {
	var s string
	var tt string

	if m.ErrorMessage == "" {

		rightPart := fmt.Sprintf("Circuit full name: %s \nCountry name: %s - Session type: %s \nTime and date of the session (UTC): %s", m.top.GranprixOffName,
			countryNameStyle.Render(m.top.CountryName), m.top.TypeSession, m.top.Date.Format("15:04:05 02-01-2006"))

		top := lipgloss.JoinHorizontal(lipgloss.Center, topLeftStyle.Render(m.top.GranprixName), topRightStyle.Render(rightPart))
		tt = ""

		if !m.tableOn {
			tt = lipgloss.JoinVertical(lipgloss.Center, topStyle.Render(top), bottomStyle.Render(fmt.Sprintf("%s waiting...\n", m.spinner.View())))
		} else {
			tt = lipgloss.JoinVertical(lipgloss.Center, topStyle.Render(top), bottomStyle.Render(m.Table.View()))
		}

		tt += helpStyle.Render(fmt.Sprintf("\n↓/↑ move\tq exit\n"))
		tt = allStyle.Render(tt)

	} else {
		tt += lipgloss.NewStyle().Foreground(lipgloss.Color("#d20f39")).Render(m.ErrorMessage)
	}

	s += lipgloss.Place(m.Terminal.widthT, m.Terminal.heightT, lipgloss.Center, lipgloss.Center, tt)

	return s
}

func CutOff(namGranPrix string) string {
	if lipgloss.Width(namGranPrix) > 50 {
		return namGranPrix[:49] + "..."
	}
	return namGranPrix
}
