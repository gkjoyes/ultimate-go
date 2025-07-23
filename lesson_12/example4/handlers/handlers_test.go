// Sample test to show how to test the execution of an internal endpoint.
// go test -run TestGetUsers -race -cpu 16
package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gkjoyes/ultimate-go/lesson_12/example4/handlers"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func init() {
	handlers.Routes()
}

// TestGetUsers testing the GET /users internal endpoint.
func TestGetUsers(t *testing.T) {
	url := "/users"
	statusCode := 200
	t.Log("Given the need to test the GetUsers endpoint.")
	{
		r := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)

		testID := 0
		t.Logf("\tTest %d:\tWhen checking %q for status code %d", testID, url, statusCode)
		{
			if w.Code != statusCode {
				t.Fatalf("\t%s\tTest %d:\tShould receive a status code of %d for the response. Received[%d].", failed, testID, statusCode, w.Code)
			}
			t.Logf("\t%s\tTest %d:\tShould receive a status code of %d for the response.", succeed, testID, statusCode)

			var users []struct {
				Name  string
				Email string
			}

			if err := json.NewDecoder(w.Body).Decode(&users); err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to decode the response.", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to decode the response.", succeed, testID)

			if len(users) == 0 {
				t.Fatalf("\t%s\tTest %d:\tShould have atleast one user in the response.", failed, testID)
			}

			if users[0].Name == "Bill" {
				t.Logf("\t%s\tTest %d:\tShould have user \"x1\" in the response.", succeed, testID)
			} else {
				t.Errorf("\t%s\tTest %d:\tShould have user \"x1\" in the response: Received[%s]", succeed, testID, users[0].Name)
			}

			if users[0].Email == "x1@gmail.com" {
				t.Logf("\t%s\tTest %d:\tShould have user \"x1@gmail.com\" in the response.", succeed, testID)
			} else {
				t.Errorf("\t%s\tTest %d:\tShould have user \"x1@gmail.com\" in the response: Received[%q]", succeed, testID, users[0].Email)
			}
		}
	}
}
