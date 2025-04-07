package tui

import (
	// "fmt"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

var (
	topLeftStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).Width(30).
			Bold(true).MarginRight(2).PaddingLeft(1).Foreground(lipgloss.Color("#cba6f7"))

	topRightStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).MarginRight(2)

	topStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).Width(100).
			Height(2).Border(lipgloss.RoundedBorder()).
			BorderLeft(false).BorderRight(false).BorderTop(false)

	bottomStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).Width(100).
			Height(20)

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

	allStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#f2cdcd"))
)

type MsgUpdate string

type Circuit struct {
	GranprixName    string
	CountryName     string
	GranprixOffName string
	Date            time.Time
	Location        string
}

type mainModel struct {
	spinner spinner.Model
	top     Circuit
	bottom  string
}

func NewModel(cir Circuit) mainModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	m := mainModel{top: cir, spinner: s}
	return m
}

func (m mainModel) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, m.spinner.Tick)
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case MsgUpdate:
		m.bottom = string(msg)
		break

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m mainModel) View() string {
	var s string

	rightPart := fmt.Sprintf("Circuit full name: %s \nCountry name: %s \nTime and date of the race (UTC): %s", m.top.GranprixOffName, m.top.CountryName, m.top.Date.Format("15:04:05 02-01-2006"))

	top := lipgloss.JoinHorizontal(lipgloss.Center, topLeftStyle.Render(m.top.GranprixName), topRightStyle.Render(rightPart))
	tt := ""
	if m.bottom == "" {
		tt = lipgloss.JoinVertical(lipgloss.Center, topStyle.Render(top), bottomStyle.Render(fmt.Sprintf("%s waiting...\n", m.spinner.View())))
	} else {
		tt = lipgloss.JoinVertical(lipgloss.Center, topStyle.Render(top), bottomStyle.Render(m.bottom))
	}
	tt += helpStyle.Render(fmt.Sprintf("\nq: exit\n"))
	tt = allStyle.Render(tt)

	width, height, err := term.GetSize(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s += lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, tt)

	return s
}
