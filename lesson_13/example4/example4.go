// Sample program to simuate io-bounded work.
package example4

import (
	"encoding/xml"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// file is mocking the XML document we expect to receive.
var file = `<?xml version="1.0" encoding="UTF-8"?>
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

func generateList(totalDocs int) []string {
	docs := make([]string, totalDocs)
	for i := 0; i < totalDocs; i++ {
		docs[i] = "test.xml"
	}
	return docs
}

func read(doc string) ([]Item, error) {
	time.Sleep(time.Millisecond) // Simulate blocking disk read.
	var d Document
	if err := xml.Unmarshal([]byte(file), &d); err != nil {
		return nil, err
	}
	return d.Channel.Items, nil
}

func find(docs []string, topic string) int {
	var found int

	for _, doc := range docs {
		items, err := read(doc)
		if err != nil {
			continue
		}

		for _, item := range items {
			if strings.Contains(item.Description, topic) {
				found++
			}
		}
	}
	return found
}

func findConcurrent(docs []string, topic string, grs int) int {
	var found int64

	var wg sync.WaitGroup
	wg.Add(grs)
	ch := make(chan string, len(docs))

	for g := 0; g < grs; g++ {
		go func() {
			defer wg.Done()

			var lFound int64
			for doc := range ch {
				items, err := read(doc)
				if err != nil {
					continue
				}

				for _, item := range items {
					if strings.Contains(item.Description, topic) {
						lFound++
					}
				}
			}
			atomic.AddInt64(&found, lFound)
		}()
	}

	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	wg.Wait()
	return int(found)
}

// func main() {
// docs := generateList(1000)
// found := find(docs, "Go")
// found := findConcurrent(docs, "Go", runtime.NumCPU())
// fmt.Println("found: ", found)
// }
