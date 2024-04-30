package main

import "fmt"
import "os"
import "os/exec"
import "time"
import "strconv"

import "github.com/charmbracelet/bubbles/timer"
import tea "github.com/charmbracelet/bubbletea"

type model struct {
	timer          timer.Model
	specialMessage string
}

func initialModel() model {
	timeoutMin := 25
	if len(os.Args) > 1 {
		arg1 := os.Args[1]
		arg1i, err := strconv.Atoi(arg1)
		if err == nil {
			timeoutMin = arg1i
		}
	}
	timeout := time.Duration(timeoutMin) * time.Minute

	return model{
		timer:          timer.NewWithInterval(timeout, time.Millisecond*100),
		specialMessage: "Press Q to quit!",
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
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.specialMessage = "Goodbye (:"
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	bs, err := exec.Command("/usr/bin/figlet", "-fslant", m.timer.View()).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(bs)

	if m.timer.Timedout() {
		s = "All done!"
	}
	s += "\n"
	s += m.specialMessage + "\n"
	return s
}

func main() {
	m := initialModel()
	if _, err := tea.NewProgram(m).Run(); err != nil {

		fmt.Println("Uh oh, we encountered an error:", err)
		os.Exit(1)
	}
}
