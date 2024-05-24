package brute

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func try(client *http.Client, u string, user string, pass string) *http.Response {
	v := url.Values{}
	v.Set("username", user)
	v.Set("password", pass)
	//finish url if needed

	if !strings.Contains(u, "login") {
		u += "/login"
	}

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

	//need to find an identifier for success
	//Identifier will be redirect, hopefully this works
	if resp.StatusCode >= 300 && resp.StatusCode <= 400 {
		return true
	}

	return false
}
