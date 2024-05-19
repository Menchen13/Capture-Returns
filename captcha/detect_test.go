package captcha

import (
	"net/http"
	"net/url"
	"testing"
)

func TestIsCaptcha(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
		want bool
	}{}

	var u string = "http://10.10.122.216"
	var client http.Client

	v := url.Values{}
	v.Set("user", "test")
	v.Set("pass", "tests2")

	r, _ := client.Get(u)
	for i := 0; i < 3; i++ {
		client.PostForm(u, v)
	}

	c, _ := client.Get(u)

	tests[0] = struct {
		name string
		args args
		want bool
	}{"noCaptcha", args{resp: r}, false}

	tests[1] = struct {
		name string
		args args
		want bool
	}{"yesCaptcha", args{resp: c}, true}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCaptcha(tt.args.resp); got != tt.want {
				t.Errorf("IsCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}
