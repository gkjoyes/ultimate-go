// Sample program to show how to create values from exported types with embedded unexporeted types.
package main

import (
	"fmt"

	"github.com/gkjoyes/ultimate-go/lesson4/exporting/example5/users"
)

func main() {
	u := users.Manager{
		Title: "Dev Manager",
	}
	u.ID = 10
	u.Name = "x1"
	fmt.Printf("User: %#v\n", u)
}
