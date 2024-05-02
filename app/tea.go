package app

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

type Selection struct {
	Label    string
	Selected bool
	Hovered  bool
}

func (s Selection) String() string {
	x := " "
	if s.Hovered && s.Selected {
		underline := color.New(color.FgHiGreen).Add(color.Underline)
		x = "x"
		return underline.Sprintf(fmt.Sprintf("[%s] %s", x, s.Label))
	}
	if s.Selected {
		x = "x"
		return color.HiGreenString(fmt.Sprintf("[%s] %s", x, s.Label))
	}

	if s.Hovered {
		underline := color.New(color.FgHiCyan).Add(color.Underline)
		x = "."

		return underline.Sprintf(fmt.Sprintf("[%s] %s", x, s.Label))
	}

	return fmt.Sprintf("[%s] %s", x, s.Label)
}

type Selections []Selection

// Select sets the selection at index i to "x" and all other selections to " ".
func (s Selections) Select(i int) Selections {
	for j := range s {
		s[j].Selected = false
	}
	s[i].Selected = true
	return s
}

func (s Selections) Hover(i int) Selections {
	for j := range s {
		s[j].Hovered = false
		if i == j {
			if s[j].Hovered {
				s[j].Hovered = false
			} else {
				s[j].Hovered = true
			}
		}
	}
	return s
}
func (s Selections) GetSelected() string {
	for _, selection := range s {
		if selection.Selected {
			return selection.Label
		}
	}
	return ""
}

func (s Selections) String() string {
	str := ""
	for _, selection := range s {

		// Render the row
		str += fmt.Sprintf("%s\t", selection.String())
	}
	return str
}

type MenuItem struct {
	Selections Selections
	Label      string
}

func (m *MenuItem) Unhover() {
	for i := range m.Selections {
		m.Selections[i].Hovered = false
	}

}
func (m *MenuItem) String(cursor string) string {
	return fmt.Sprintf("%s:\n %s\t%s",  m.Label,cursor,m.Selections.String())
}
type MetaballsViewModel struct {
	Menu       []MenuItem
	HorzCursor int // which to-do list item our cursor is pointing at
	VertCursor int // which to-do item our cursor is pointing at


}

var (
	METABALL_SELECTIONS = map[string]*Selections{
		"Ball Size": &Selections{
			Selection{Label: "Small", Selected: true},
			Selection{Label: "Medium", Selected: false},
			Selection{Label: "Large", Selected: false},
		},
		"Ball Speed": &Selections{
			Selection{Label: "Slow", Selected: true},
			Selection{Label: "Medium", Selected: false},
			Selection{Label: "Fast", Selected: false},
		},

		"Screen Size": &Selections{
			Selection{Label: "Small", Selected: true},
			Selection{Label: "Medium", Selected: false},
			Selection{Label: "Large", Selected: false},
		},
		"Resolution": &Selections{
			Selection{Label: "Low", Selected: true},
			Selection{Label: "Medium", Selected: false},
			Selection{Label: "High", Selected: false},
		},
		"FPS": &Selections{
			Selection{Label: "30", Selected: false},
			Selection{Label: "45", Selected: true},
			Selection{Label: "60", Selected: false},
		},

		"Ball Count": &Selections{
			Selection{Label: "4", Selected: true},
			Selection{Label: "8", Selected: false},
			Selection{Label: "20", Selected: false},
		},

	}
)

func initialModel() MetaballsViewModel {
	return MetaballsViewModel{
		// Our to-do list is a grocery list
		Menu: []MenuItem{
			MenuItem{
				Label:      "Ball Size",
				Selections: *METABALL_SELECTIONS["Ball Size"],
			},
			MenuItem{
				Label:      "Ball Speed",
				Selections: *METABALL_SELECTIONS["Ball Speed"],
			},

			MenuItem{
				Label:      "Screen Size",
				Selections: *METABALL_SELECTIONS["Screen Size"],
			},
			MenuItem{
				Label:      "Resolution",
				Selections: *METABALL_SELECTIONS["Resolution"],
			},
			MenuItem{
				Label:      "FPS",
				Selections: *METABALL_SELECTIONS["FPS"],
			},

			MenuItem{
				Label:      "Ball Count",
				Selections: *METABALL_SELECTIONS["Ball Count"],
			},
		},
		// The cursor, which indicates the item selected
		HorzCursor: 0,
		VertCursor: 0,


	}
}

func (m MetaballsViewModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m MetaballsViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.Menu[m.VertCursor].Selections.Hover(m.HorzCursor)
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "right", "l":

			// If the cursor is at the top of the list, don't move it up
			if m.HorzCursor < len(m.Menu[m.VertCursor].Selections)-1 {
				m.HorzCursor++
				m.Menu[m.VertCursor].Selections.Hover(m.HorzCursor)
			}

		case "left", "j":

			// If the cursor is at the bottom of the list, don't move it down
			if m.HorzCursor > 0 {
				m.HorzCursor--
				m.Menu[m.VertCursor].Selections.Hover(m.HorzCursor)
			}

		case "up", "i":

			if m.VertCursor > 0 {
				m.Menu[m.VertCursor].Unhover()
				m.VertCursor--
				m.Menu[m.VertCursor].Selections.Hover(m.HorzCursor)

			}

		case "down", "k":

			if m.VertCursor < len(m.Menu)-1 {
				m.Menu[m.VertCursor].Unhover()
				m.VertCursor++
				m.Menu[m.VertCursor].Selections.Hover(m.HorzCursor)
			}

		case "enter":
			m.Menu[m.VertCursor].Selections.Select(m.HorzCursor)
		case "q":
			return m, tea.Quit
		case "ctrl+c":
			os.Exit(0)
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m MetaballsViewModel) View() string {
	// The header
	s := "Select your MetaBall settings\n\n"

	// Iterate over our choices
	for vi, item := range m.Menu {
		// Is the cursor pointing at this choice?
		vCursor := " " // no cursor
		if m.VertCursor == vi {
			vCursor = ">" // cursor!
		}
		s += item.String(vCursor)

		// Add a newline after each item
		s += "\n\n"
	}

	// The footer
	s += "\n\nPress q to confirm, or ctrl+c to exit.\n"

	// Send the UI for rendering
	return s
}