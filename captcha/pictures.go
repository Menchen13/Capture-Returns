package captcha

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/png"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/otiai10/gosseract/v2"
)

// takes in the base64encoded image string and returns the name of the shape as a string
// NOT IMPLEMENTET YET!!
func shape(b64encoded string) (string, error) {
	//decode b64encoded into reader
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64encoded))

	img, format, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("Format: ", format, "Color model: ", img.ColorModel())

	return "", nil
}

// takes in base64-legal sting containing the term image and returns the solved term.
func term(base64encoded string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	//getting testing image string from Responses
	str, err := base64.StdEncoding.DecodeString(base64encoded)
	if err != nil {
		return "", err
	}

	client.SetWhitelist("0123456789+*-/?=")
	client.Trim = true
	err = client.SetImageFromBytes(str)
	if err != nil {
		return "", err
	}

	term, err := client.Text()

	//cut of "=?" for eval
	term = term[:len(term)-2]

	result, err := eval(term)
	if err != nil {
		return "", err
	}

	// convert the result to string using fmt.Sprint
	var a = fmt.Sprint(result)
	return a, nil
}

func eval(s string) (int, error) {
	eval, err := govaluate.NewEvaluableExpression(s)
	if err != nil {
		return 0, err
	}

	result, err := eval.Eval(nil)
	if err != nil {
		return 0, err
	}

	return int(result.(float64)), nil

}
