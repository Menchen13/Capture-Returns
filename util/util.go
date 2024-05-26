package util

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/fs"
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

func B64encodedToFile(b64encoded string) error {
	arr, err := base64.RawStdEncoding.DecodeString(b64encoded)
	if err != nil {
		return err
	}
	fs.FileMode
	err = os.WriteFile("square.png", arr)
	if err != nil {
		return err
	}
}
