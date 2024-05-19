package captcha

import (
	"net/http"
	"strings"
)

func IsCaptcha(resp *http.Response) bool {

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

func isShape(resp *http.Response) bool {
	arr := make([]byte, 330)

	resp.Body.Read(arr)
	defer resp.Body.Close()

	if strings.Contains(string(arr), "shape") {
		return true
	}
	return false
}
