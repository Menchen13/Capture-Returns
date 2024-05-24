package captcha

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/otiai10/gosseract/v2"
	"github.com/stretchr/testify/assert"
)

func TestTerm(t *testing.T) {
	t.Run("Multiplication with tmpfile", func(t *testing.T) {
		t.Skip("Skippin test due to lack of necessity")
		var b64img string = "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"

		//getting testing image string from Responses
		str, err := base64.StdEncoding.DecodeString(b64img)
		if err != nil {
			t.Fatal("non nil error when decoding image string! ", err)
		}

		tmpfile, err4 := os.CreateTemp("", "ocrfile")
		//cleanup for tmpfile
		defer func() {
			tmpfile.Close()
			os.Remove(tmpfile.Name())
		}()

		if err4 != nil {
			assert.Error(t, err4, "error when creating tmp file", err4)
		}

		_, err5 := tmpfile.Write(str)
		if err5 != nil {
			assert.Error(t, err5, "error when writing to tmp file", err5)
		}

		client := gosseract.NewClient()
		defer client.Close()

		client.SetLanguage()

		err2 := client.SetWhitelist("0123456789+*-/?=")
		if err2 != nil {
			assert.Error(t, err2, "error when setting whitelist", err2)
		}
		client.Trim = true

		client.SetImage(tmpfile.Name())

		term, erro := client.Text()
		if erro != nil {
			t.Fatal("non nil error when calling .Text() function", erro)
		}

		assert.EqualValuesf(t, "308*26=?", term, "Non matching Terms")
	})
	t.Run("Multiplicantion without tmpfile", func(t *testing.T) {
		var b64img string = "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"

		//getting testing image string from Responses
		str, err := base64.StdEncoding.DecodeString(b64img)
		if err != nil {
			t.Fatal("non nil error when decoding image string! ", err)
		}

		client := gosseract.NewClient()
		defer client.Close()

		err2 := client.SetWhitelist("0123456789+*-/?=")
		if err2 != nil {
			assert.Error(t, err2, "error when setting whitelist", err2)
		}
		client.Trim = true

		client.SetImageFromBytes(str)

		term, erro := client.Text()
		if erro != nil {
			t.Fatal("non nil error when calling .Text() function", erro)
		}
		assert.Equal(t, "308*26=?", term)
		//cut of "=?" for eval
		term = term[:len(term)-2]
		assert.Equal(t, "308*26", term)

	})
	t.Run("Actual func", func(t *testing.T) {
		var b64img string = "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"

		result, err := term(b64img)
		if err != nil {
			t.Error("Error thrown by Term: ", err)
		}
		assert.Equal(t, 8008, result)

	})
}
