package display

type ProjectRow struct {
	ID          uint64 `header:"id"`
	Visibility  string `header:"visibility"`
	Name        string `header:"name"`
	Description string `header:"description"`
}
