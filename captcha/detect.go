package captcha

import (
	"fmt"
	"net/http"
	"strings"
)

func IsCaptcha(resp *http.Response) bool {

	//220 characters in body before lable and image
	arr := make([]byte, 220)
	resp.Body.Read(arr)
	defer resp.Body.Close()

	//check for captcha in arr
	fmt.Println(string(arr))
	if strings.Contains(string(arr), "captcha") {
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
