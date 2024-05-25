package util

import (
	"fmt"
	"testing"
)

func TestRespFromFile(t *testing.T) {
	resp := RespFromFile("..\\Responses\\circle.html")
	if resp == nil {
		t.Fatal("resp is nil")
	}

	fmt.Println(resp.Status)
}
