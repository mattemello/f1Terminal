package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	topStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Left).Width(100).
			Height(2).BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69"))

	bottomStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Left).Width(100).
			Height(15).BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69"))

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type mainModel struct {
	top    int
	bottom int
}

func NewModel() mainModel {
	m := mainModel{}
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
	s += lipgloss.JoinVertical(lipgloss.Top, topStyle.Render(fmt.Sprintf("Hello\n")), bottomStyle.Render(fmt.Sprintf("world\n")))
	s += helpStyle.Render(fmt.Sprintf("\nq: exit\n"))
	return s
}
