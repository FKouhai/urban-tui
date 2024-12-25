package model
import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/FKouhai/urban-cli/urbanapi"
	tea "github.com/charmbracelet/bubbletea"
	"fmt"
)
type errMsg error
type Model struct {
	Spinner spinner.Model
	TextInput textinput.Model
	Err error
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
			s := urbanapi.Run(value)
			return m, tea.Println(s)
		}
	case errMsg:
		m.Err = msg
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}
func (m Model) View() string {
	s := fmt.Sprintf("def -> %s",m.TextInput.View())
	return s
}
