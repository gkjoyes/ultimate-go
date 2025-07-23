// Sample test to show how to mock an HTTP GET call internally.
package example3

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// feed is mocking the XML document we expect to receive.
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/golang/go</description>
    <link>https://go.dev//</link>
    <item>
        <pubDate>Sun, 13 Mar 2025 15:04:00 +0000</pubDate>
        <title>Multi Paradigm language Mechanics</title>
        <description>Go is a multi-paradigm language.</description>
        <link>https://go.dev/ref/spec</link>
    </item>
</channel>
</rss>`

// Document defines the fields associated with the RSS document.
type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel defines the fields associated with the channel tag in the RSS document.
type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	Items       []Item   `xml:"item"`
}

// Item defines the fields associated with the item tag in the RSS document.
type Item struct {
	XMLName     xml.Name `xml:"item"`
	PubDate     string   `xml:"pubDate"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

// mockServer returns a pointer to a server to handle the mock get call.
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload validates the http Get function can download content and the content can be unmarshalled and clean.
func TestDownload(t *testing.T) {
	statusCode := http.StatusOK
	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen checking %q for status code %d", testID, server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to make the Get call: %v", failed, testID, err)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to make the Get call.", succeed, testID)
			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t%s\tTest %d:\tShould receive a %d status code: %v", failed, testID, statusCode, resp.StatusCode)
			}
			t.Logf("\t%s\tTest %d:\tShould receive a %d status code.", succeed, testID, statusCode)

			var d Document
			if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to unmarshal the response : %v", failed, testID, err)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to unmarshal the response.", succeed, testID)

			if len(d.Channel.Items) == 1 {
				t.Logf("\t%s\tTest %d:\tShould have 1 item in the feed.", succeed, testID)
			} else {
				t.Errorf("\t%s\tTest %d:\tShould have 1 item in the feed: %d", failed, testID, len(d.Channel.Items))
			}
		}
	}
}
