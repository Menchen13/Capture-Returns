package brute

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_try(t *testing.T) {
	r := try(http.DefaultClient, "http://10.10.122.216/login", "test1", "test2")

	var arr []byte = make([]byte, r.ContentLength)
	fmt.Println(string(arr))
}
