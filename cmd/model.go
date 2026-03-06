package main

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Help  key.Binding
	Quit  key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right}, // first column
		{k.Help, k.Quit},                // second column
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type styles struct {
	app           lipgloss.Style
	title         lipgloss.Style
	statusMessage lipgloss.Style
}

type Task struct {
	status      status
	title       string
	description string
}

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

// Main Tea Model
type Model struct {
	list          []list.Model
	loaded        bool
	focused       status
	help          help.Model
	keys          keyMap
	styles        styles
	width, height int
	err           error
}

func New() *Model {
	return &Model{focused: todo, help: help.New(), keys: keys}
}

func (m *Model) updateListProperties() {
	// Update list size.
	h, v := m.styles.app.GetFrameSize()
	for i := range m.list {
		m.list[i].SetSize(m.width-h, m.height-v)
		m.list[i].Styles.Title = m.styles.title

	}
}

// "*" modifies the model directly
func (m *Model) initList(width int, height int) {
	colWidth := (width / 3) - 2
	colHeight := height - 2 // reserve space for help bar

	todoList := list.New([]list.Item{}, list.NewDefaultDelegate(), colWidth, colHeight)
	inProgressList := list.New([]list.Item{}, list.NewDefaultDelegate(), colWidth, colHeight)
	doneList := list.New([]list.Item{}, list.NewDefaultDelegate(), colWidth, colHeight)

	m.list = []list.Model{todoList, inProgressList, doneList}

	for i := range m.list {
		m.list[i].SetShowHelp(false)
		switch i {
		case int(todo):
			m.list[i].Title = "To Do"
			m.list[i].SetItems([]list.Item{
				Task{status: todo, title: "buy milk", description: "strawberry milk"},
				Task{status: todo, title: "do something", description: "do something again"},
				Task{status: todo, title: "buy fruits", description: "banana"},
			})

		case int(inProgress):
			m.list[i].Title = "In Progress"
			m.list[i].SetItems([]list.Item{
				Task{status: inProgress, title: "do something", description: "what what"},
			})
			m.list[i].Select(-1)

		case int(done):
			m.list[i].Title = "Done"
			m.list[i].SetItems([]list.Item{
				Task{status: done, title: "finished task", description: "this task is finished"},
			})
			m.list[i].Select(-1)
		}
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
