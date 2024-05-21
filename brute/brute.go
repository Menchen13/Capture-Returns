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
// NOT IMPLEMENTET YET
func Orca(Url string, user string, pass string) bool {
	resp := try(http.DefaultClient, Url, user, pass)

	//still need to ajust lenght of read in
	arr := make([]byte, 100)
	resp.Body.Read(arr)

	//need to find an identifier for success

	return false
}
