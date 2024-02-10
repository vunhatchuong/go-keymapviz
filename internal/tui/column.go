package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type column struct {
	title        string
	selectedItem int
	focus        bool
	list         list.Model
	width        int
	height       int
}

func newColumn(title string) column {
	var focus bool
	if title == "keyboards" {
		focus = true
	}
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	defaultList.SetShowHelp(false)
	return column{focus: focus, title: title, list: defaultList}
}

func (c *column) Focus() {
	c.focus = true
}

func (c *column) Blur() {
	c.focus = false
}

func (c *column) Focused() bool {
	return c.focus
}

func (c *column) getNext() int {
	items := c.list.Items()
	numberOfItems := len(items)
	if c.list.Index()+2 > numberOfItems {
		return c.selectedItem
	}
	c.selectedItem = c.list.Index() + 1
	return c.selectedItem
}

func (c *column) getPrev() int {
	if c.list.Index()-1 < 0 {
		return c.selectedItem
	}
	c.selectedItem = c.list.Index() - 1
	return c.selectedItem
}

func (c column) Init() tea.Cmd {
	return nil
}

func (c column) View() string {
	return c.getStyle().Render(c.list.View())
}

func (c column) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.setSize(msg.Width/2-2, msg.Height/3)
		c.list.SetSize(msg.Width/2-2, msg.Height/3)

	case tea.KeyMsg:
	}
	c.list, cmd = c.list.Update(msg)
	return c, cmd
}

func (c *column) setSize(w, h int) {
	c.width = w
	c.height = h
}

func (c *column) getStyle() lipgloss.Style {
	if c.Focused() {
		return lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Width(c.width).
			Height(c.height)
	}
	return lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.HiddenBorder()).
		Width(c.width).
		Height(c.height)
}
