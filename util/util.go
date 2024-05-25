package util

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// returns the content of file at path as a *http.Response
// jumps to /workdir before trying to open the file
// therefore path to file should be from workdir
func RespFromFile(path string) *http.Response {
	os.Chdir("/workdir")
	arr, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error reading file", err)
		return nil
	}
	reader := bytes.NewReader(arr)
	bufreader := bufio.NewReaderSize(reader, len(arr))
	resp, err := http.ReadResponse(bufreader, nil)
	if err != nil {
		fmt.Println("error in ReadResponse function", err)
		return nil
	}
	return resp
}
