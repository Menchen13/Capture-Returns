package brute

import (
	"fmt"
	"net/http"
	"net/url"
)

func try(client *http.Client, u string, user string, pass string) *http.Response {
	v := url.Values{}
	v.Set("user", user)
	v.Set("pass", pass)

	resp, err := client.PostForm(u, v)
	if err != nil {
		fmt.Println("Error trying user pass combination", err)
	}
	return resp
}
