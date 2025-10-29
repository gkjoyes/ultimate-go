// Sample program to show how one need to be careful when appending to a slice when you have a reference to an element.
package main

import "fmt"

type user struct {
	likes int
}

func main() {
	// Declare a slice of 3 users.
	users := make([]user, 3)

	// Share the user at index 1.
	user1 := &users[1]

	// Add a like for the user that was shared.
	user1.likes++

	displayLikes(users)

	fmt.Println("---------------------------")

	// Add a new user.
	users = append(users, user{})

	// Add another like for the user that was shared.
	user1.likes++

	// Notice the last like has not been recorded.
	displayLikes(users)
}

// displayLikes displays the number of likes for all users.
func displayLikes(users []user) {
	for i := range users {
		fmt.Printf("User: [%d]\tLikes: [%d]\n", i, users[i].likes)
	}
}
