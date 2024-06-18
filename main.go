package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	techList map[string][]string
	cursor   int
	selected map[string]string
	listInd  []string
	index    int
	done     bool
}

func initialModel() model {
	return model{
		techList: map[string][]string{
			"os":       {"linux"},
			"server":   {"apache", "node.js", "unicorn", "tomcat", "nginix", "kestrel", "cowboy"},
			"database": {"mySQL", "mongoDB", "postgreSQL", "microsoftSQL"},
			"language": {"PHP", "javascript", "ruby", "java", "python", "c#", "elixir", "go"},
		},
		cursor:   0,
		selected: make(map[string]string),
		listInd:  []string{"os", "server", "database", "language"},
		index:    0,
		done:     false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	if m.done {
		output := "Your selected tech-stack is: \n\n"
		for _, key := range m.listInd {
			output += fmt.Sprintf("%s --> %s\n", key, m.selected[key])
		}
		output += "\npress ESC to quit"
		return output
	}
	var options string
	options += "Welcome to tech-stack composer for your project\n\n"
	options += "The tech-stack contains the following components\n-> OS\n-> Server\n-> Database\n-> Language\n"
	switch m.listInd[m.index] {
	case "os":
		options += "\n"
		options += "OS\nThe OS is fixed to LINUX\n\npress enter to move next step"

	case "server":
		options += fmt.Sprintf("\n")
		options += "Server\nYou can select any one of the servers given below\n\n"

		for i, choice := range m.techList[m.listInd[m.index]] {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			options += fmt.Sprintf("%s %s\n", cursor, choice)
		}
		options += "\nPress Enter to select an item and move to the next list\n"
		options += "Press ESC to quit\n"

	case "database":
		options += fmt.Sprintf("\n")
		options += "Database\nYou can select any one of the databases given below\n\n"
		for i, choice := range m.techList[m.listInd[m.index]] {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			options += fmt.Sprintf("%s %s\n", cursor, choice)
		}
		options += "\nPress Enter to select an item and move to the next list\n"
		options += "Press ESC to quit\n"

	case "language":
		options += fmt.Sprintf("\n")
		options += "\nLanguage\nYou can select any one of the languages given below\n\n"
		for i, choice := range m.techList[m.listInd[m.index]] {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			options += fmt.Sprintf("%s %s\n", cursor, choice)
		}
		options += "\nPress Enter to select an item and move to the next list\n"
		options += "Press ESC to quit\n"
	}

	return options
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := m.listInd[m.index]
		switch msg.String() {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.techList[key])-1 {
				m.cursor++
			}
		case "enter":
			m.selected[key] = m.techList[key][m.cursor]
			if m.index < len(m.listInd)-1 {
				m.index++
				m.cursor = 0
			} else {
				m.done = true
			}
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

func main() {
	program := tea.NewProgram(initialModel())

	if _, err := program.Run(); err != nil {
		fmt.Printf("Error occurs:\n%v", err)
		os.Exit(1)
	}
}
