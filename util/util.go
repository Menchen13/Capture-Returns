package util

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// returns the content of file at path as a *http.Response
func RespFromFile(path string) *http.Response {
	arr, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error reading file")
		return nil
	}
	reader := bytes.NewReader(arr)
	bufreader := bufio.NewReaderSize(reader, len(arr))
	resp, err := http.ReadResponse(bufreader, nil)
	if err != nil {
		fmt.Println("error in ReadResponse function")
		return nil
	}
	return resp
}
