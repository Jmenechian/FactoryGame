package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	choices []string
	cursor int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices: []string{"Option 1", "Option 2", "Option 3"},

		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

			case "ctrl+c", "q":
				return m, tea.Quit
			
			case "up", "k" :
				if m.cursor > 0 {
					m.cursor--
				}
			
			case "down":
				if m.cursor < 2 {
					m.cursor++
				}

			case "enter", "space":
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	s := "What do you want to do?\n\n"

	for i, choice := range m.choices {
	
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return tea.NewView(s)
}