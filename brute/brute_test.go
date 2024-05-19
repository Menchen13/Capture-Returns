package brute

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_try(t *testing.T) {
	r := Try(http.DefaultClient, "http://10.10.122.216/", "test1", "test2")

	var arr []byte = make([]byte, r.ContentLength)
	r.Body.Read(arr)
	defer r.Body.Close()
	fmt.Println(string(arr))
}
