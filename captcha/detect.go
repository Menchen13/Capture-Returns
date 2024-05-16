package captcha

import (
	"net/http"
	"strings"
)

func IsCaptcha(resp *http.Response, url string) bool {

	//220 characters in body before lable and image
	arr := make([]byte, 220)
	resp.Body.Read(arr)

	//check for captcha in arr
	if strings.Contains(string(arr), "captcha") {
		return true
	}

	return false
}
