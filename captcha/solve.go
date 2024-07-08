package captcha

import (
	"fmt"
	"net/http"
	"net/url"
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

		fmt.Println("SOLVER LOOP: ", i) //debug

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
		v := url.Values{}
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
// needs to be redone and retestet
func getImage(resp *http.Response) string {
	//redoo
	return ""
}
