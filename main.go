// main package
package main

import (
	"flag"
	"fmt"
)

func main() {

	url := flag.String("u", "", "Url to attack")
	userfile := flag.String("l", "", "Path to file with usernames")
	passfile := flag.String("p", "", "Path to file with passwords")
	flag.Parse()
	if *url == "" {
		fmt.Println("No Url specified")
		flag.PrintDefaults()
		return
	} else if *userfile == "" {
		fmt.Println("No userfile specified")
		flag.PrintDefaults()
		return
	} else if *passfile == "" {
		fmt.Println("No passfile specified")
		flag.PrintDefaults()
	}

	//client := captcha.Sync(*url)
	//defer client.Close()

}
