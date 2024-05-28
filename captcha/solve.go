package captcha

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

		//get b64 encoded image sting from response
		var img = getImage(resp)
		fmt.Println("got image: ", i) //debug

		//check for type of captcha and call responding solve function
		var answer string
		if isShape(resp) {
			fmt.Println("image: ", i, "is shape") //debug
			answer, err = shape(img)
		} else {
			fmt.Println("image:", i, "is term") //debug
			answer, err = term(img)
		}

		if err != nil {
			panic(err)
		}
		//create url Value and add answer to it
		var v url.Values
		v.Add("captcha", answer)
		//send answer using PostForm func and get the response as input for next iteration
		resp, err = http.PostForm(u, v)
		if err != nil {
			panic(err)
		}
		fmt.Println("captcha: ", i, "solved") //debug
	}
	return
}

// takes in a http Response and returns the b64 encoded image string
func getImage(resp *http.Response) string {
	//fuck it. Whole body it is.

	var arr = make([]byte, resp.ContentLength)
	resp.Body.Read(arr)
	var str = string(arr)

	_, str, a := strings.Cut(str, "src")
	if !a {
		panic("Couldnt cut out 'src' ")
	}

	//cuts of everything until the b64 string
	_, str, a = strings.Cut(str, ",")
	if !a {
		panic("unable to Cut to ','")
	}
	fmt.Println("str after ',' split: ", str) //debug

	//cuts of everything after the b64 string
	str, _, a = strings.Cut(str, "\"")
	if !a {
		panic("unable to cut to \"")
	}
	fmt.Println("str after cutting to \": ", str) //debug

	return str
}
