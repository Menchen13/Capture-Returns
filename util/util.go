package util

import (
	"bufio"
	"bytes"
	"encoding/base64"
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

// takes in a b64 encoded string, containing a png image and returns a file with the image
// it is the callers responsibility to remove the file when it is nolonger needed
func B64ToFile(b64encoded string) (*os.File, error) {
	arr, err := base64.RawStdEncoding.DecodeString(b64encoded)
	if err != nil {
		//remove later
		if len(b64encoded) >= 79171 {
			fmt.Println("79171: ", []byte(b64encoded)[79171])
		} else {
			fmt.Println("22095: ", string([]byte(b64encoded)[22095]))
		}
		//22095
		return nil, err
	}
	//default perms rw-r--r-- are: 0664
	tmp, err := os.CreateTemp("", "")
	if err != nil {
		return nil, err
	}

	defer tmp.Close()
	_, err = tmp.Write(arr)

	if err != nil {
		return nil, err
	}

	return tmp, nil
}
