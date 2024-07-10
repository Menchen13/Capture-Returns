package captcha

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// returns true if the captcha is shape based
func isShape(resp *http.Response) bool {
	//read in response Body
	BytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("Error Reading in Body: %w", err))
	}
	defer resp.Body.Close()

	// Create a new goquery document from the file
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(BytesBody))
	if err != nil {
		panic(fmt.Errorf("Error reading in BytesBody: %w", err))
	}

	var shape bool = false
	// Find the label element that contains the description
	doc.Find("label").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		fmt.Println("found label", text) //debug
		if strings.Contains(text, "Describe the shape below") {
			shape = true
		}
	})

	return shape
}
