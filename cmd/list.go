package cmd

import (
	"fmt"
	"os"

	"github.com/filipponova/sm/internal"
	"github.com/spf13/cobra"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type instanceItem struct {
	inst internal.Instance
}

func (i instanceItem) Title() string       { return i.inst.Name }
func (i instanceItem) Description() string { return i.inst.ID }
func (i instanceItem) FilterValue() string { return i.inst.Name }

type tuiModel struct {
	list     list.Model
	selected instanceItem
	quitting bool
}

func (m *tuiModel) Init() tea.Cmd {
	return nil
}

func (m *tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if item, ok := m.list.SelectedItem().(instanceItem); ok {
				m.selected = item
				m.quitting = true
				return m, tea.Quit
			}
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *tuiModel) View() string {
	if m.quitting && m.selected.inst.ID != "" {
		return "Connecting to instance: " + m.selected.inst.Name + " (" + m.selected.inst.ID + ")\n"
	}
	return m.list.View()
}

func listAndConnectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available EC2 instances and connect to one using AWS Session Manager.",
		Run: func(cmd *cobra.Command, args []string) {
			instances, err := internal.GetEC2Instances(region, profile)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			if len(instances) == 0 {
				fmt.Println("No instances available for session manager.")
				os.Exit(0)
			}
			items := make([]list.Item, len(instances))
			for i, inst := range instances {
				items[i] = instanceItem{inst}
			}
			// Adjust the list height to show all instances at once
			height := len(items)
			if height < 5 {
				height = 5 // minimum height for UX
			}
			l := list.New(items, list.NewDefaultDelegate(), 50, height+2)
			l.Title = "Select an EC2 instance"
			m := &tuiModel{list: l}
			p := tea.NewProgram(m)
			finalModel, err := p.Run()
			if err != nil {
				fmt.Println("Error starting TUI:", err)
				os.Exit(1)
			}
			selected := finalModel.(*tuiModel).selected
			if selected.inst.ID != "" {
				fmt.Printf("Starting session on instance %s ...\n", selected.inst.ID)
				internal.StartSession(region, profile, selected.inst.ID)
			}
		},
	}
	return cmd
}
