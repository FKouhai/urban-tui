package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)
type RData struct {
	Data []struct {
		Meaning string `json:"meaning"`
		Example string `json:"example"`
	}
}

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
func run(message string) string{
	endpoint := "https://unofficialurbandictionaryapi.com/api/search?term="
	log.Println("fetching definition")
	var d RData
	out, err := connect(endpoint,message); if err != nil {
		log.Println(err)
	}
	d1,err := decode(out,d); if err!= nil {
		log.Println(err)
	}
  return fmt.Sprintf("Meaning %s ",d1.Data[0].Meaning)
}

func decode(data []byte, dStruct RData) (*RData,error){
	err := json.Unmarshal(data, &dStruct); if err != nil {
		log.Fatalln("Unable to unmarshal data", err)
	}
	return &dStruct, nil
}

func connect(endpoint string, term string) ([]byte,error){
	var client http.Client
	search := endpoint + term
	req, err := http.NewRequest("GET", search, nil); if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(req); if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body); if err != nil {
		log.Fatalln(err)
		return nil,err
	}
	return body,nil
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
			s := run(value)
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
