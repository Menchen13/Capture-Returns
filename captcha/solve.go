package captcha

import (
	"net/http"
	"net/url"
	"strings"
)

// Solves all captchas until another attempt is possible
// NOT IMPLEMENTET YET!!!
func Solver(u string) {
	for i := 0; i < 3; i++ {
		resp, err := http.Get(u)
		if err != nil {
			panic(err)
		}
		//get b64 encoded image sting from response
		var img = getImage(resp)

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
		//What is the server variable name?
		var v url.Values
		v.Add("?", answer)
		//need to add the Postform for aswering captcha.

	}
	return
}

// takes in a http Response and returns the b64 encoded image string
// NOT IMPLEMENTET YET!!!
func getImage(resp *http.Response) string {
	//fuck it. Whole body it is.

	var arr = make([]byte, resp.ContentLength)
	resp.Body.Read(arr)
	//about 1400 characters till body starts
	var str = string(arr[1400:])

	strings.C

	return ""
}
