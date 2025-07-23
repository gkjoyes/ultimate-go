// Sample to show how to write a basic example test.
// go test -run ExampleGetUsers
package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func ExampleGetUsers() {
	r := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	var users []struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(w.Body).Decode(&users); err != nil {
		log.Fatal("ERROR: ", err)
	}
	fmt.Println(users)
	// Output:
	// [{x1 x1@gmail.com}]
}
