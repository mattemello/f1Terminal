package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	topLeftStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).Width(30).
			Bold(true).MarginRight(2).PaddingLeft(1).Foreground(lipgloss.Color("#cba6f7"))

	topRightStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).MarginRight(2)

	topStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).Width(100).
			Height(2).BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69"))

	bottomStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).Width(100).
			Height(15).BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69"))

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type Circuit struct {
	GranprixName    string
	CountryName     string
	GranprixOffName string
	Date            time.Time
	Location        string
}

type mainModel struct {
	top    Circuit
	bottom int
}

func NewModel(cir Circuit) mainModel {
	m := mainModel{top: cir}
	return m
}

func (m mainModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m mainModel) View() string {
	var s string

	rightPart := fmt.Sprintf("Circuit full name: %s \nCountry name: %s \nTime and date of the race (UTC): %s", m.top.GranprixOffName, m.top.CountryName, m.top.Date.Format("15:04:05 02-01-2006"))

	top := lipgloss.JoinHorizontal(lipgloss.Center, topLeftStyle.Render(m.top.GranprixName), topRightStyle.Render(rightPart))
	tt := lipgloss.JoinVertical(lipgloss.Center, topStyle.Render(top), bottomStyle.Render(fmt.Sprintf("world\n")))
	tt += helpStyle.Render(fmt.Sprintf("\nq: exit\n"))

	// todo: center the place
	s += lipgloss.Place(100, 17, lipgloss.Center, lipgloss.Center, tt)

	return s
}
