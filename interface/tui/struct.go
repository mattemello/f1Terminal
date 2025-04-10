package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
)

type MsgError struct {
	Message string
}

type MsgUpdateCiruit Circuit

type MsgUpdateTable struct {
	SessionOn bool
	Table     []table.Row
}

type Circuit struct {
	GranprixName    string
	CountryName     string
	GranprixOffName string
	Date            time.Time
	Location        string
	TypeSession     string
}

type Terminal struct {
	widthT  int
	heightT int
}

type mainModel struct {
	spinner      spinner.Model
	top          Circuit
	Table        table.Model
	Terminal     Terminal
	ErrorMessage string
	tableOn      bool
	sessionOn    bool
	altScreen    bool
}
