package main

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width, msg.Height)
		m.help.SetWidth(msg.Width)
		m.loaded = true
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Left):
			m.list[m.focused].Select(-1)
			if m.focused == 0 {
				m.focused = 2
			} else {
				m.focused--
			}

		case key.Matches(msg, m.keys.Right):
			m.list[m.focused].Select(-1)
			if m.focused == 2 {
				m.focused = 0
			} else {
				m.focused++
			}

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	if m.loaded {
		m.list[m.focused], cmd = m.list[m.focused].Update(msg)
	}
	return m, cmd
}
