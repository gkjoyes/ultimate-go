// package users provide support for user management.
package users

// user represents information about a user.
type user struct {
	ID   int
	Name string
}

// Manager represents information about a manager.
type Manager struct {
	Title string
	user
}
