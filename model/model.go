package model

import (
	//"fmt"

	"fmt"

	"github.com/FKouhai/urban-cli/urbanapi"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
type errMsg error
type Styles struct {
	BorderColor lipgloss.Color
	InputField lipgloss.Style
	OutputField lipgloss.Style
}
func DefaultStyLes() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().Foreground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	s.OutputField = lipgloss.NewStyle().Foreground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	return s
}
type Model struct {
	Styles *Styles
	Height int
	Width int
	Spinner spinner.Model
	TextInput textinput.Model
	Err error
	Done bool
	Definition string
	Key string
	Example string
}
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			value := m.TextInput.Value()
			s,e := urbanapi.Run(value)
			m.Key = value
			m.Definition = s
			m.Example = e
			m.Done = true
			return m,tea.EnterAltScreen
		case tea.KeyCtrlN:
			m.Done = false
			return m,tea.ClearScreen
		}
	case errMsg:
		m.Err = msg
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.Done {
	return	lipgloss.Place(
			m.Width,
			m.Height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Left,
				fmt.Sprintf("%s", m.Key),
				m.Styles.OutputField.Render(m.Definition),
				fmt.Sprintf("Example"),
				m.Styles.OutputField.Render(m.Example)),
				)
				//fmt.Sprintf("%s -> %s\nexample -> %s", m.Key, m.Definition, m.Example)),
			//)
	}
	return lipgloss.Place(
		m.Width,
		m.Height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			"Enter the slang you want to know the definition of:",
			m.Styles.InputField.Render(m.TextInput.View()),),
		)
}
