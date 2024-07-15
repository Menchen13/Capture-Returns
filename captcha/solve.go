package captcha

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// Solves all captchas until another attempt is possible
// NOT TESTET YET
func Solver(u string) {
	resp, err := http.Get(u)
	if err != nil {
		panic(err)

	}

	//do everything 3 times as captchas come in batches of 3
	for i := 0; i < 3; i++ {

		Body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic("error when reading in resp.Body")
		}
		resp.Body.Close()

		//get b64 encoded image sting from response
		var img = getImage(Body)

		//check for type of captcha and call responding solve function
		var answer string
		if isShape(Body) {
			answer, err = shape(img)
		} else {
			answer, err = term(img)
		}

		if err != nil {
			panic(err)
		}
		//create url Value and add answer to it
		v := url.Values{}
		v.Add("captcha", answer)
		//send answer using PostForm func and get the response as input for next iteration
		resp, err = http.PostForm(u, v)
		if err != nil {
			panic(err)
		}
	}

	Body, _ := io.ReadAll(resp.Body)
	if strings.Contains(string(Body), "captchas") {
		fmt.Println("Solver seems to have left a captcha open. Restarting solver!")
		v := url.Values{}
		v.Add("captcha", "tranglange") //obviously wrong answer to reset captcha counter, so the loop of 3 is correct
		http.PostForm(u, v)
		Solver(u) //calling solver function again
	}
	return
}

// takes in a http Response and returns the b64 encoded image string
func getImage(arr []byte) string {

	doc, err := html.Parse(strings.NewReader(string(arr)))
	if err != nil {
		panic(err)
	}

	var imageSrc string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					imageSrc = a.Val
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if imageSrc == "" {
		fmt.Println(string(arr)) //debug
		panic("image src not found")
	}
	_, imageSrc, _ = strings.Cut(imageSrc, ",")

	return imageSrc

}
