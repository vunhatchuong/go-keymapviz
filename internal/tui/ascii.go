package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type columnViewPort struct {
	focus    bool
	ready    bool
	viewport viewport.Model
	width    int
	height   int
	content  string
	help     help.Model
}

func newViewPort(art string) columnViewPort {
	help := help.New()
	help.ShowAll = true

	return columnViewPort{content: art, help: help}
}

func (c *columnViewPort) Focus() {
	c.focus = true
}

func (c *columnViewPort) Blur() {
	c.focus = false
}

func (c *columnViewPort) Focused() bool {
	return c.focus
}

func (c *columnViewPort) setSize(w, h int) {
	c.width = w
	c.height = h
}

func (m columnViewPort) Init() tea.Cmd {
	return nil
}

func (m columnViewPort) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		footerHeight := lipgloss.Height(m.helpView())
		m.setSize(msg.Width-2, msg.Height/2)
		if !m.ready {
			m.viewport = viewport.New(msg.Width-2, (msg.Height/2)-footerHeight)
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width - 2
			m.viewport.Height = (msg.Height / 2) - footerHeight
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m columnViewPort) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return m.getStyle().Render(m.viewport.View())
}

func (m columnViewPort) helpView() string {
	return m.help.View(keys)
}

func (c *columnViewPort) getStyle() lipgloss.Style {
	if c.Focused() {
		return lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Width(c.width).
			Height(c.height).
			Align(lipgloss.Center, lipgloss.Center)
	}
	return lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.HiddenBorder()).
		Width(c.width).
		Height(c.height).
		Align(lipgloss.Center, lipgloss.Center)
}
