package main

import "fmt"
import "os"
import "time"

import "github.com/charmbracelet/bubbles/timer"
import tea "github.com/charmbracelet/bubbletea"

const timeout = time.Second * 5

type model struct {
	timer timer.Model
}

func initialModel() model {
	return model{
		timer: timer.NewWithInterval(timeout, time.Millisecond),
	}
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	s := m.timer.View()

	if m.timer.Timedout() {
		s = "All done!"
	}
	s += "\n"
	return s
}

func main() {
	m := initialModel()
	if _, err := tea.NewProgram(m).Run(); err != nil {

		fmt.Println("Uh oh, we encountered an error:", err)
		os.Exit(1)
	}
}
