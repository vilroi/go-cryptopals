package base64

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vilroi/go-cryptopals/hex"
)

var TestVals = map[string]string{
	"light work.": "bGlnaHQgd29yay4=",
	"light work":  "bGlnaHQgd29yaw==",
	"light wor":   "bGlnaHQgd29y",
	"light wo":    "bGlnaHQgd28=",
	"light w":     "bGlnaHQgdw==",
}

func TestBase64Decode(t *testing.T) {
	s := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	ans := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	bytes, err := hex.EncodeStr(s)
	if err != nil {
		t.Fatal(err)
	}
	res := Decode(bytes)
	assert.Equal(t, ans, res, "failed")

	for k, v := range TestVals {
		res = Decode([]byte(k))
		assert.Equal(t, v, res, "failed")
	}
}
