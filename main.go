// main package
package main

import (
	"flag"
)

func main() {

	url := flag.String("u", "", "Url to attack")
	flag.Parse()
	if *url == "" {
		flag.PrintDefaults()
		return
	}

}
