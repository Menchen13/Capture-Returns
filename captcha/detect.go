package captcha

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// returns true if the captcha is shape based
func isShape(resp *http.Response) bool {
	//read in response Body
	BytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("error reading in body: %w", err))
	}
	defer resp.Body.Close()

	// doc, err := html.Parse(bytes.NewReader(BytesBody))

	// if err != nil {
	// 	panic(fmt.Errorf("error parsing: '%w'", err))
	// }

	// var shape bool = false
	// var f func(*html.Node)
	// f = func(n *html.Node) {
	// 	// fmt.Println("Type: ", n.Type, "Data: ", n.Data, "Namespace: ", n.Namespace, "Attributes: ", n.Attr)
	// 	if n.Type == 1 && n.Data == "Describe the shape below (circle, square, or triangle)" {

	// 		shape = true
	// 	}
	// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 		f(c)
	// 	}
	// }
	// f(doc)

	// return shape

	if strings.Contains(string(BytesBody), "circle, square, or triangle") {
		return true
	}
	fmt.Println(string(BytesBody)) //debug

	return false
}
