package captcha

import (
	"fmt"
	"net/http"
	"strings"
)

func Sync(url string) *http.Client {
	client := http.Client{}
	for i := 0; i < 3; i++ {
		resp, err := client.Get(url)

		if err != nil {
			panic(err)
		} else if resp.StatusCode != 200 {

			fmt.Println("Non 200 Statuscode on try", i, ": ", resp.StatusCode)
		}
	}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	//220 characters in body before lable and image
	arr := make([]byte, 220)
	resp.Body.Read(arr)

	//check for captcha in arr
	if strings.Contains(string(arr), "captcha") {
		fmt.Println("Captcha synced!")
	}

	return &client
}
