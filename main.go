// main package
package main

import (
	"flag"
	"fmt"
	"net/http"
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
		return
	}

	//check if url is reachable

	_, err := http.Get(*url)
	if err != nil {
		fmt.Println("Error reaching url!", err)
	}

	//client := captcha.Sync(*url)
	//defer client.Close()

}
