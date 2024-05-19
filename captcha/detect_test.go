package captcha

import (
	"net/http"
	"net/url"
	"testing"
)

func TestIsCaptcha(t *testing.T) {
	//tests successfull skipping this now as it is no longer functional
	t.SkipNow()
	//variable settup (god this is horrific)
	var u string = "http://10.10.128.201"
	var client http.Client

	v := url.Values{}
	v.Set("user", "test")
	v.Set("pass", "tests2")

	r, _ := client.Get(u)
	//send 3 tries to get sever to initiate captcha
	for i := 0; i < 3; i++ {
		//brute.Try(&client, u, "user", "pass")
	}
	c, _ := client.Get(u)

	//defining test cases
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "noCaptcha",
			args: args{resp: r},
			want: false,
		},
		{
			name: "yesCaptcha",
			args: args{resp: c},
			want: true,
		},
	}

	//actuall tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCaptcha(tt.args.resp); got != tt.want {
				t.Errorf("IsCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isShape(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isShape(tt.args.resp); got != tt.want {
				t.Errorf("isShape() = %v, want %v", got, tt.want)
			}
		})
	}
}
