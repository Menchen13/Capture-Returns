package captcha

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// returns true if the captcha is shape based
func isShape(resp *http.Response) bool {
	//read in response Body
	BytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("Error Reading in Body: %w", err))
	}
	defer resp.Body.Close()

	doc, err := html.Parse(bytes.NewReader(BytesBody))
	if err != nil {
		panic(fmt.Errorf("Error parsing HTML: %v\n", err))

	}

	// Traverse the HTML nodes to find the label
	var labelText string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "label" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode && strings.Contains(c.Data, "Describe the shape below") {
					labelText = c.Data
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if labelText != "" {
		return false
	} else {
		return true
	}
}
