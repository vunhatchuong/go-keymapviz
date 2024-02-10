package tui

type Item struct {
	Name, Desc string
}

func NewItem(title, description string) Item {
	return Item{Name: title, Desc: description}
}

func (i Item) Title() string       { return i.Name }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.Name }
