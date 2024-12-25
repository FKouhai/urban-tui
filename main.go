package main

import (
	"log"
	"github.com/FKouhai/urban-cli/model"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func initialModel() model.Model {
	ti := textinput.New()
	s := spinner.New()
	s.Spinner = spinner.Dot
	ti.Placeholder = "based"
	ti.Focus()
	ti.CharLimit = 144
	return model.Model {
		TextInput: ti,
		Spinner: s,
		Err: nil,
	}
}

//TODO -> need to improve the visuals(lipgloss)
//TODO -> usage of huh
//TODO -> if want to make a new search add in the case statement ctrl + n for a new search
//TODO -> need to use a prompt for the word to search
//TODO -> create a readme using VHS
