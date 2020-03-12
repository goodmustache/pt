package display

type UserRow struct {
	ID       uint64 `header:"user id"`
	Username string `header:"username"`
	Name     string `header:"name"`
	Email    string `header:"email"`
}
