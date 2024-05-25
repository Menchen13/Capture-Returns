package captcha

import (
	"net/http"
	"net/url"
	"strings"
)

// Solves all captchas until another attempt is possible
// TODOO implement shape() and finish post form
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
			answer, err = shape(img)
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
func getImage(resp *http.Response) string {
	//fuck it. Whole body it is.

	var arr = make([]byte, resp.ContentLength)
	resp.Body.Read(arr)
	//about 1400 characters till body starts
	var str = string(arr[1400:])

	_, str, a := strings.Cut(str, "src")
	if !a {
		panic("Couldnt cut cout to response 1")
	}
	//cuts of everything until the b64 string
	str = str[strings.Index(str, ",")+1:]
	//cuts of everything after the b64 string
	str = str[:strings.IndexByte(str, byte('"'))]

	return str
}
