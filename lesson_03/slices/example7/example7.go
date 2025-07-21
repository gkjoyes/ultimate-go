// Sample program to show how to declare and use variadic function.
package main

import "fmt"

// user represent infromation about a user.
type user struct {
	id   int
	name string
}

func main() {
	u1 := user{
		id:   1,
		name: "x1",
	}

	u2 := user{
		id:   2,
		name: "x2",
	}

	// display the both user values.
	display(u1, u2)

	users := []user{
		{3, "x3"},
		{4, "x4"},
	}

	// display the all user values from the slice.
	display(users...)

	// udpate second user.
	change(users...)

	// display the all user values again.
	display(users...)
}

// display can accept and display multiple values of user types.
func display(users ...user) {
	fmt.Printf("\n")
	for _, u := range users {
		fmt.Printf("User: [%+v]\n", u)
	}
}

// change shows how the backing array is shared.
func change(users ...user) {
	if len(users) < 2 {
		return
	}
	users[1] = user{id: 22, name: "x22"}
}
