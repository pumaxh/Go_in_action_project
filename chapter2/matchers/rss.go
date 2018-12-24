// Matcher for searching rss feeds
package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/Go_in_action_project/chapter2/search"
)

type (
	// item defines the fields associated with the item tag
	// int the rss document
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"PubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guild"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image defaines the fields associated with the image tag
	// in the rss document.
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel defines the fields associated with the channel tag
	// in the rss document
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"titel"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"image"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument defines the fields associated with the rss document
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel `xml:"channel"`
	}
)
// rssMatcher implements the Matcher interface
type rssMatcher struct {}

// init registers the matcher with the programe
func init()  {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// retrieve performs a HTTP Get request for the rss feed and decodes
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == " " {
		return nil, errors.New("No rss feed URL provided")
	}

	// Retrieve the rss feed document from the web
	rsp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// Close the response once we return from the function
	defer rsp.Body.Close()

	// Check the status code for a 200 so we know we have received a
	// proper response
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", rsp.StatusCode)
	}

	// Decode the rss feed document into our struct type
	// We don't need to check for errors, the caller can do this
	var document rssDocument
	err = xml.NewDecoder(rsp.Body).Decode(&document)
	return &document, err
}