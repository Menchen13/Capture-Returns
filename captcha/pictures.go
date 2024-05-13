package captcha

import (
	"encoding/base64"

	"github.com/Knetic/govaluate"
	"github.com/otiai10/gosseract/v2"
)

// takes in base64-legal sting containing the term image and returns the solved term.
func Term(base64encoded string) (int, error) {
	client := gosseract.NewClient()
	defer client.Close()

	//getting testing image string from Responses
	str, err := base64.StdEncoding.DecodeString(base64encoded)
	if err != nil {
		return 0, err
	}

	client.SetWhitelist("0123456789+*-/?=")
	client.Trim = true
	err = client.SetImageFromBytes(str)
	if err != nil {
		return 0, err
	}

	term, err := client.Text()

	//cut of "=?" for eval
	term = term[:len(term)-2]

	result, err := eval(term)
	if err != nil {
		return 0, err
	}

	return result, nil
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
