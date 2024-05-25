package captcha

import (
	"net/http"
	"strings"
)

// returns true if the response contains a captcha
func IsCaptcha(u string) bool {
	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}

	//1550 characters in body before lable and image
	arr := make([]byte, 1550)
	resp.Body.Read(arr)
	defer resp.Body.Close()

	//check for captcha in arr
	//only check last little bit
	if strings.Contains(string(arr[1325:]), "captcha") {
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
