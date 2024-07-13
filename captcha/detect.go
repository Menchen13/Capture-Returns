package captcha

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

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

	var f func(*html.Node)
	f = func(n *html.Node) {
		//
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {

			f(c)
		}
	}
	f(doc)

	return false
}
