// Sample test to show how to write a basic sub unit table test.

// go test -v
// go test -run TestDownload/statusok -v
// go test -run TestDownload/statusnotfound -v
// go test -run TestParallelizeDownload -v
package example5_test

import (
	"net/http"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// TestDownload validates the http Get function can download content and handles different status conditions properly.
func TestDownload(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		statusCode int
	}{
		{
			name:       "statusok",
			url:        "https://github.com/gkjoyes/ultimate-go",
			statusCode: 200,
		},
		{
			name:       "statusnotfound",
			url:        "https://github.com/gkjoyes/ultimate-go/tests",
			statusCode: 404,
		},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for id, test := range tests {
			sub := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen checking %q for status code %d", id, test.url, test.statusCode)
				{
					resp, err := http.Get(test.url)
					if err != nil {
						t.Fatalf("\t%s\tTest %d:\tShould be able to make the Get call: %v", failed, id, err)
					}
					t.Logf("\t%s\tTest %d:\tShould be able to make the Get call.", succeed, id)

					defer resp.Body.Close()

					if resp.StatusCode == test.statusCode {
						t.Logf("\t%s\tTest %d:\tShould receive a %d status code.", succeed, id, test.statusCode)
					} else {
						t.Errorf("\t%s\tTest %d:\tShould receive a %d status code: %v", failed, id, test.statusCode, resp.StatusCode)
					}
				}
			}
			t.Run(test.name, sub)
		}
	}
}

// TestParallelizeDownload validates the http Get function can download content and handles different status conditions properly.
// but runs the tests in parallel.
func TestParallelizeDownload(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		statusCode int
	}{
		{
			name:       "statusok",
			url:        "https://github.com/gkjoyes/ultimate-go",
			statusCode: 200,
		},
		{
			name:       "statusnotfound",
			url:        "https://github.com/gkjoyes/ultimate-go/tests",
			statusCode: 404,
		},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for id, test := range tests {
			sub := func(t *testing.T) {
				t.Parallel()

				t.Logf("\tTest %d:\tWhen checking %q for status code %d", id, test.url, test.statusCode)
				{
					resp, err := http.Get(test.url)
					if err != nil {
						t.Fatalf("\t%s\tTest %d:\tShould be able to make the Get call: %v", failed, id, err)
					}
					t.Logf("\t%s\tTest %d:\tShould be able to make the Get call.", succeed, id)

					defer resp.Body.Close()

					if resp.StatusCode == test.statusCode {
						t.Logf("\t%s\tTest %d:\tShould receive a %d status code.", succeed, id, test.statusCode)
					} else {
						t.Errorf("\t%s\tTest %d:\tShould receive a %d status code: %v", failed, id, test.statusCode, resp.StatusCode)
					}
				}
			}
			t.Run(test.name, sub)
		}
	}
}
