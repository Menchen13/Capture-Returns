package captcha

import (
	"Menchen13/Capture-Returns/util"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/otiai10/gosseract/v2"
	"gocv.io/x/gocv"
)

// takes in the base64encoded image string and returns the name of the shape as a string
// this whole thing is one fat chat-gpt grab
func shape(b64encoded string) (string, error) {
	file, err := util.B64ToFile(b64encoded)
	if err != nil {
		return "", err
	}
	//delete tmp file at the end of func
	defer os.Remove(file.Name())

	//read in file
	imgMat := gocv.IMRead(file.Name(), gocv.IMReadGrayScale)
	if imgMat.Empty() {
		return "", errors.New("empty image mat")
	}
	defer imgMat.Close()

	//remove noise from file (probably not needed in my case but gpt is GOD)
	blured := gocv.NewMat()
	gocv.GaussianBlur(imgMat, &blured, image.Pt(5, 5), 0, 0, gocv.BorderDefault)
	defer blured.Close()

	//find edges
	edges := gocv.NewMat()
	gocv.Canny(blured, &edges, 50, 150)
	defer edges.Close()

	//find contours (whatever the fuck that is)
	contours := gocv.FindContours(edges, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	defer contours.Close()

	contour := contours.At(0)

	peri := gocv.ArcLength(contour, true)
	approx := gocv.ApproxPolyDP(contour, 0.04*peri, true)

	defer approx.Close()

	switch approx.Size() {
	case 3:
		return "triangle", nil
	case 4:
		return "square", nil
	default:
		return "circle", nil
	}
}

// takes in base64-encoded sting containing the term image and returns the solved term.
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
		return "", fmt.Errorf("SetImageFromBytes(): %w", err)
	}

	term, err := client.Text()

	if err != nil || len(term) == 0 {
		return "", fmt.Errorf("Text(): %w", err)
	}

	//cut of "=?" for eval
	term = strings.TrimRight(term, "=?")

	result, err := eval(term)
	if err != nil {
		fmt.Println(term)
		return "", fmt.Errorf("eval(): %w", err)
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
