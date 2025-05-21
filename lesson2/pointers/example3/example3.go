// Sample program to show the mechanics of escape analysis.
package main

// user represents a user in the system.
type user struct {
	name  string
	email string
}

func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("Address u1: [", &u1, "]")
	println("Address u2: [", u2, "]")
}

// createUserV1 creates a user value and passed a copy back to the caller.
//
//go:noinline
func createUserV1() user {
	u := user{
		name:  "george",
		email: "george@gmail.com",
	}
	return u
}

// createUserV2 creates a user value and shares the value with the caller.
//
//go:noinline
func createUserV2() *user {
	u := user{
		name:  "george",
		email: "george@gmail.com",
	}
	return &u
}
