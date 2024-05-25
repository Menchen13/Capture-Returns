package util

import (
	"fmt"
	"testing"
)

func TestRespFromFile(t *testing.T) {
	//path needs to match layout of dev container.
	resp := RespFromFile("Responses/circle.html")
	if resp == nil {
		t.Fatal("resp is nil")
	}

	fmt.Println(resp.Status)
}
