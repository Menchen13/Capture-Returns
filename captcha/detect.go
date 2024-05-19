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
	fmt.Println(arr)
	if strings.Contains(string(arr), "captcha") {
		return true
	}

	return false
}
