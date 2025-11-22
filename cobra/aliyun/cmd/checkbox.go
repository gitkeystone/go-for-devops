/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func updateCheckbox(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices) {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else if m.cursor != len(m.choices) {
				m.selected[m.cursor] = struct{}{}
			} else {
				m.cursor = 0
				for i := range m.selected {
					m.Language = append(m.Language, m.choices[i])
				}
				return m, tea.Quit
			}
		}

	}

	return m, nil
}

func checkboxView(m model) string {
	var b strings.Builder

	b.WriteString("Which language should we configure?\n\n")

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		fmt.Fprintf(&b, "%s [%s] %s\n", cursor, checked, choice)
	}

	button := &blurredButton
	if m.cursor == len(m.choices) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString("\nPress q to quit.\n")

	return b.String()
}
