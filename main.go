package main

import "fmt"
import "os"
import "os/exec"
import "time"

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	tick      *time.Ticker
	startTime time.Time
}

func initialModel() model {
	return model{
		tick:      time.NewTicker(time.Millisecond * 30),
		startTime: time.Now(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "The time is\n"

	s += fmt.Sprintf("I'm ticking at: ", m.tick.Sub(startTime))
}

func main() {
	//	go func() {
	//		for {
	//			mainLoop(<-tick.C, startTime)
	//		}
	//	}()
	time.Sleep(time.Minute)
}

func mainLoop(tickTime time.Time, startTime time.Time) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("I'm ticking at: ", tickTime.Sub(startTime))
}
