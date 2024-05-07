// main package
package main

import (
	"Menchen13/Capture-Returns/captcha"
	"fmt"
)

func main() {

	_, err := captcha.Term()

	fmt.Println(err)

}
