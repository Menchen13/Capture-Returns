package brute

import (
	"net/http"
	"net/url"
)

func try(client *http.Client, u *string, user string, pass string) bool {
	v := url.Values{}
	v.Set("user", user)
	v.Set("pass", pass)

	client.PostForm(*u, v)
	return false
}
