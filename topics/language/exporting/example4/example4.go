// Sample program to show how unexported fields from an exported struct type can't be accessed directly.
package main

import (
	"fmt"

	"github.com/gkjoyes/ultimate-go/topics/exporting/example4/users"
)

func main() {
	// Create a value of type User from the users package.
	u := users.User{
		ID:   10,
		Name: "x1",
		// password: "xxxx",  NOTE: can't access this field directly.
	}

	fmt.Printf("%#v\n", u)
}
