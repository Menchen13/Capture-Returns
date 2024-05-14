// main package
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
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
		fmt.Println("Error reaching url! ", err)
		return
	}

	//read in user and pass file
	user, err := os.ReadFile(*userfile)
	if err != nil {
		fmt.Println("Error reading userfile! ", err)
		return
	}

	pass, err := os.ReadFile(*passfile)
	if err != nil {
		fmt.Println("Error reading passfile! ", err)
		return
	}

	//turn user and pass file into slices for easier iteration
	var userSlice []string = strings.FieldsFunc(string(user), linebreak)
	var passSlice []string = strings.FieldsFunc(string(pass), linebreak)

	if userSlice == nil {
		fmt.Println("Error while slicing userfile! ")
	}

	if passSlice == nil {
		fmt.Println("Error while slicing passfile!")
	}

	//main bruteforce loop
	/*
		for i, v := range userSlice {
			for n, k := range passSlice {
				if captcha.test {
					captacha.solve()
				} else {
					if try(v, k) {
						fmt.Println("Success! ", v, ":", k)
						return
					}
				}
			}
		}
	*/

}

func try(user string, pass string)

func linebreak(r rune) bool {
	if r == '\n' || r == '\r' {
		return true
	} else {
		return false
	}
}
