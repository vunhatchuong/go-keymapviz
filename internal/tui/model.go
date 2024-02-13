package tui

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vunhatchuong/go-keymapviz/keyboards"
)

type columnName int

const (
	KEYBOARD columnName = iota
	LAYOUT
	ART
)

func (s columnName) getNext() columnName {
	if s == ART {
		return KEYBOARD
	}

	return s + 1
}

func (s columnName) getPrev() columnName {
	if s == KEYBOARD {
		return ART
	}

	return s - 1
}

type Model struct {
	focused  columnName
	cols     []column
	viewport columnViewPort
}

func NewModel() *Model {
	return &Model{focused: KEYBOARD}
}

func (m *Model) InitList() {
	m.cols = []column{
		newColumn("keyboards"),
		newColumn("layouts"),
	}

	kbs := keyboards.Keyboards

	m.cols[KEYBOARD].list.Title = "Keyboards"
	kbList := make([]list.Item, 0, len(kbs))

	for kb := range kbs {
		kbList = append(kbList, Item{Name: kb})
	}

	m.cols[KEYBOARD].list.SetItems(kbList)

	m.cols[LAYOUT].list.Title = "Layouts"
	selectedKb := m.SelectedItemName(KEYBOARD)
	layoutList := make([]list.Item, 0, len(kbs[selectedKb]))

	for layout := range kbs[selectedKb] {
		layoutList = append(layoutList, Item{Name: layout})
	}

	m.cols[LAYOUT].list.SetItems(layoutList)

	// Init ascii
	kb := m.SelectedItemName(KEYBOARD)
	layout := m.SelectedItemName(LAYOUT)
	layouts := keyboards.Keyboards[kb]
	m.viewport = newViewPort(layouts[layout])
}

func (m *Model) changeFocusedNext(next bool) {
	if m.focused == ART {
		m.viewport.Blur()
	} else {
		m.cols[m.focused].Blur()
	}

	if next {
		m.focused = m.focused.getNext()
	} else {
		m.focused = m.focused.getPrev()
	}

	if m.focused == ART {
		m.viewport.Focus()
	} else {
		m.cols[m.focused].Focus()
	}
}

func (m Model) SelectedItemName(column columnName) string {
	col := m.cols[column]

	item, ok := col.list.Items()[col.selectedItem].(Item)
	if !ok {
		log.Panic("Not an item")
	}

	return item.Title()
}

func (m Model) populateLayouts() {
	kb := m.SelectedItemName(KEYBOARD)
	layouts := keyboards.Keyboards[kb]
	itemList := make([]list.Item, 0, len(layouts))

	for layout := range layouts {
		itemList = append(itemList, Item{Name: layout})
	}

	m.cols[LAYOUT].list.SetItems(itemList)
}

func (m *Model) populateAscii() {
	kb := m.SelectedItemName(KEYBOARD)
	layout := m.SelectedItemName(LAYOUT)
	layouts := keyboards.Keyboards[kb]
	m.viewport.viewport.SetContent(layouts[layout])
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		for i := 0; i < len(m.cols); i++ {
			var res tea.Model
			res, cmd = m.cols[i].Update(msg)
			m.cols[i] = res.(column)

			cmds = append(cmds, cmd)
		}

		var res tea.Model
		res, cmd = m.viewport.Update(msg)
		m.viewport = res.(columnViewPort)

		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, keys.Left):
			m.changeFocusedNext(false)
		case key.Matches(msg, keys.Right):
			m.changeFocusedNext(true)

		case key.Matches(msg, keys.Up):
			m.cols[m.focused].getPrev()

			if m.focused == KEYBOARD {
				m.populateLayouts()
			}

			m.populateAscii()

		case key.Matches(msg, keys.Down):
			m.cols[m.focused].getNext()

			if m.focused == KEYBOARD {
				m.populateLayouts()
			}

			m.populateAscii()
		}
	}

	var res tea.Model

	if m.focused == ART {
		res, cmd = m.viewport.Update(msg)
		if _, ok := res.(columnViewPort); ok {
			m.viewport = res.(columnViewPort)

			cmds = append(cmds, cmd)
		}

		return m, tea.Batch(cmds...)
	}

	res, cmd = m.cols[m.focused].Update(msg)

	if _, ok := res.(column); ok {
		m.cols[m.focused] = res.(column)

		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	ui := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.cols[0].View(),
		m.cols[1].View(),
	)
	ui += lipgloss.JoinVertical(
		lipgloss.Bottom,
		`\n`,
		m.viewport.View(),
	)
	ui += lipgloss.JoinVertical(
		lipgloss.Right,
		`\n`,
		m.viewport.helpView(),
	)

	return ui
}
