// Sample program to show how to use error variables to help the caller determine the exact error being returned.
package main

import (
	"errors"
	"fmt"
)

var (
	// ErrBadRequest is returned when there are problems witht the request.
	ErrBadRequest = errors.New("Bad Request")
	// ErrPageMoved is returned when a 302/301 is returned.
	ErrPageMoved = errors.New("Page Moved")
)

func main() {
	if err := webCall(true); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request Occurred")
			return
		case ErrPageMoved:
			fmt.Println("The Page Moved")
			return
		default:
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Life is Good")
}

// webCall performs a web operation.
func webCall(b bool) error {
	if b {
		return ErrBadRequest
	}
	return ErrPageMoved
}
