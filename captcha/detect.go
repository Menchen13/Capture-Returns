package captcha

import (
	"net/http"
	"strings"
)

// returns true if the response from url "u" contains a captcha
func IsCaptcha(u string) bool {
	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}

	//whole body read in
	arr := make([]byte, resp.ContentLength)
	resp.Body.Read(arr)
	defer resp.Body.Close()

	//check everything
	if strings.Contains(string(arr), "captcha") {
		return true
	}

	return false
}

// returns true if the captcha is shape based
func isShape(resp *http.Response) bool {
	//about 1650 bytes till image
	arr := make([]byte, 1700)

	resp.Body.Read(arr)
	defer resp.Body.Close()

	//only check in last little bit of body
	if strings.Contains(string(arr[1500:]), "shape") {
		return true
	}
	return false
}
