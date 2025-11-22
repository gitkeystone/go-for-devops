/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "配置中心",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := tea.NewProgram(initialModel()).Run()
		if err != nil {
			fmt.Printf("could not start program: %s\n", err)
			os.Exit(1)
		}

		bytes, err := json.MarshalIndent(m, "", "    ")
		if err != nil {
			panic(err)
		}

		// fmt.Println(string(bytes))

		if err := os.WriteFile(profile, bytes, os.ModePerm); err != nil {
			panic(err)
		}
	},
}

var profile string

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	configureCmd.Flags().StringVarP(&profile, "profile", "p", "config.json", "配置文件")
}

const (
	inputboxState int = iota
	radioState
	checkboxState
)

type model struct {
	state int

	// inputbox
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode

	ID  string `json:"id,omitempty"`
	Key string `json:"key,omitempty"`

	// radio
	cursor int
	choice string

	ChinaRegion string `json:"region,omitempty"`

	// checkbox
	choices  []string
	selected map[int]struct{}

	Language []string `json:"language,omitempty"`
}

func initialModel() model {
	// inputbox
	m := model{
		state:  inputboxState,
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

	// radio
	// checkbox

	m.choices = []string{"zh", "en", "jp"}
	m.selected = make(map[int]struct{})

	return m
}

func (m model) Init() tea.Cmd {
	if m.state == 0 {
		return textinput.Blink
	} else {
		return nil
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case inputboxState:
		return updateInputbox(m, msg)
	case radioState:
		return updateRadio(m, msg)
	case checkboxState:
		return updateCheckbox(m, msg)
	}

	return m, nil
}

func (m model) View() string {
	switch m.state {
	case inputboxState:
		return inputboxView(m)
	case radioState:
		return radioView(m)
	case checkboxState:
		return checkboxView(m)
	}

	return ""
}
