package captcha

import (
	"net/http"
	"net/url"
)

// Solves all captchas until another attempt is possible
// NOT IMPLEMENTET YET!!!
func Solver(u string) {
	for i := 0; i < 3; i++ {
		resp, err := http.Get(u)
		if err != nil {
			panic(err)
		}
		var img string
		//need to extract image string from resp

		//check for type of captcha and call responding solve function
		var answer string
		if isShape(resp) {
			answer = shape(img)
		} else {
			answer, err = term(img)
		}

		if err != nil {
			panic(err)
		}

		var v url.Values
		v.Add("?", answer)
		//need to add the Postform for aswering captcha.
		//What is the server variable name?

	}
	return
}

// takes in a http Response and returns the b64 encoded image string
// NOT IMPLEMENTET YET!!!
func getImage(resp *http.Response) string {
	return ""
}
