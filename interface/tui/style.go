package tui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	width        = 100
	heightBottom = 23
	heightTop    = 2
	heightTotal  = 32

	topLeftStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).Width(width / 4).
			Bold(true).MarginRight(2).PaddingLeft(1).Foreground(lipgloss.Color("#cba6f7"))

	topRightStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).MarginRight(2)

	topStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).Width(width).
			Height(heightTop).Border(lipgloss.RoundedBorder()).
			BorderLeft(false).BorderRight(false).BorderTop(false)

	bottomStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Top).Width(width).
			Height(heightBottom)

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

	countryNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#91d7e3"))

	allStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#f2cdcd"))

	columsPos = []table.Column{
		{Title: "Position", Width: 8},
		{Title: "First Name", Width: 9},
		{Title: "Last Name", Width: 10},
		{Title: "N° Driver", Width: 9},
		{Title: "Team Name", Width: 15},
	}

	columsSess = []table.Column{
		{Title: "Position", Width: 8},
		{Title: "First Name", Width: 9},
		{Title: "Last Name", Width: 10},
		{Title: "N° Driver", Width: 9},
		{Title: "Gap Leader", Width: 9},
		{Title: "Gap Head", Width: 9},
		{Title: "Team Name", Width: 10},
	}
)
