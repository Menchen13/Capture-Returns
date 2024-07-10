package testing

import (
	"Menchen13/Capture-Returns/util"
	"net/http"
	"testing"
)

func TestHtmlParse(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Testing",
			args: args{resp: util.RespFromFile("Responses/square.html")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HtmlParse(tt.args.resp); got != tt.want {
				t.Errorf("HtmlParse() = %v, want %v", got, tt.want)
			}
		})
	}
}
