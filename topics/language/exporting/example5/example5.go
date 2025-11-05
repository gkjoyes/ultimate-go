// Sample program to show how to create values from exported types with embedded unexported types.
package main

import (
	"fmt"

	"github.com/gkjoyes/ultimate-go/topics/language/exporting/example5/users"
)

func main() {
	u := users.Manager{
		Title: "Dev Manager",
	}
	u.ID = 10
	u.Name = "x1"
	fmt.Printf("User: %#v\n", u)
}
