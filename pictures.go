package main

import (
	"encoding/base64"
	"fmt"

	"github.com/otiai10/gosseract"
)

func string() ([]byte, error) {
	client := gosseract.NewClient()
	defer client.Close()

	//getting testing image string from Responses
	str, err := base64.StdEncoding.DecodeString("")
	if err != nil {
		return []byte{}, err
	}

	client.SetWhitelist("0123456789+*-/")
	client.Trim = true
	client.SetImageFromBytes(str)
	fmt.Println(client.Text())

	return []byte{}, err
}
