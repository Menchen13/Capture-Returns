package util

import (
	"bufio"
	"bytes"
	"net/http"
	"os"
)

// returns the content of file at path as a *http.Response
func RespFromFile(path string) *http.Response {
	arr, _ := os.ReadFile(path)
	reader := bytes.NewReader(arr)
	bufreader := bufio.NewReader(reader)
	resp, _ := http.ReadResponse(bufreader, nil)
	return resp
}
