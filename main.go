package main

import (
	"log"
	"github.com/FKouhai/urban-cli/urban-api"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err error
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "based"
	ti.Focus()
	ti.CharLimit = 144
	return model {
		textInput: ti,
		err: nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			value := m.textInput.Value()
			s := urbanapi.Run(value)
			return m, tea.Println(s)
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}
func (m model) View() string {
	return m.textInput.View()
}
