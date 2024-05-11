package captcha

import (
	"fmt"
	"net/http"
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

	//220 characters in body before layble and image
	arr := make([]byte, 220)
	resp.Body.Read(arr)
	//check for captcha in arr

	return &client
}
