/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/spf13/cobra"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "阿里云配置中心",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := tea.NewProgram(initialModel()).Run()
		if err != nil {
			fmt.Printf("could not start program: %s\n", err)
			os.Exit(1)
		}

		// Assert the final tea.Model to our local model and print the choice.
		if m, ok := m.(model); ok && m.choice != "" {
			fmt.Printf("\n---\nYou chose %s!\n", m.choice)
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	configureCmd.Flags().StringVarP(&profile, "profile", "", "profile.json", "配置文件")
}

var profile string

// Bubble Tea
var choices = []string{"cn-shanghai", "cn-hangzhou"}

var (
	// TextInputs
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))

	// Choices
	keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	ticksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
)

type modelState int

const (
	TextInputs modelState = iota
	Choice
	Choices
	dotChar = " • "
)

type model struct {
	state    modelState
	Quitting bool

	// TextInputs
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode

	// Choice
	cursor int
	choice string

	// Choices

}

func initialModel() model {
	// TextInputs
	m := model{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32
		t.Width = 20

		switch i {
		case 0:
			t.Placeholder = "Access Secret ID"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Access Secret Key"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		m.inputs[i] = t
	}

	// Choice
	// Choices

	return m
}

func (m model) Init() tea.Cmd {
	m.state = TextInputs
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	switch m.state {
	case TextInputs:
		return updateTextInputs(m, msg)
	case Choice:
		return updateChoice(m, msg)
	case Choices:
		return updateChoices(m, msg)
	}

	return m, nil
}

func updateTextInputs(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				m.state = Choice
				return m, nil
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func updateChoice(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.cursor++
			if m.cursor > 1 {
				m.cursor = 1
			}
		case "k", "up":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = 0
			}
		case "enter":
            m.choice = choices[m.cursor]
			return m, tea.Quit
		}
	}

	return m, nil
}

func updateChoices(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m model) View() string {
	if m.Quitting {
		return "\n  See you later!\n\n"
	}

	switch m.state {
	case TextInputs:
		return textInputsView(m)
	case Choice:
		return choiceView(m)
	case Choices:
		return choicesView(m)
	}

	return ""
}

func textInputsView(m model) string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

func choiceView(m model) string {
	c := m.cursor

	tpl := "Choose a Region\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s",
		checkbox(choices[c], c == 0),
		checkbox(choices[c], c == 1),
	)

	return fmt.Sprintf(tpl, choices)
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}

func choicesView(m model) string {
	return ""
}

// Utils

// Generate a blend of colors.
func makeRampStyles(colorA, colorB string, steps float64) (s []lipgloss.Style) {
	cA, _ := colorful.Hex(colorA)
	cB, _ := colorful.Hex(colorB)

	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, lipgloss.NewStyle().Foreground(lipgloss.Color(colorToHex(c))))
	}
	return
}

// Convert a colorful.Color to a hexadecimal format.
func colorToHex(c colorful.Color) string {
	return fmt.Sprintf("#%s%s%s", colorFloatToHex(c.R), colorFloatToHex(c.G), colorFloatToHex(c.B))
}

// Helper function for converting colors to hex. Assumes a value between 0 and
// 1.
func colorFloatToHex(f float64) (s string) {
	s = strconv.FormatInt(int64(f*255), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return
}
