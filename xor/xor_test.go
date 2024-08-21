package xor

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/vilroi/goutils/assert"
)

func TestXorStrings(t *testing.T) {
	text := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	//text := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal!"		// false
	key := "ICE"
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	xored := XorStrings(text, key)
	res := hex.EncodeToString(xored)

	assert.AssertEq(expected, res)
	fmt.Println("OK")
}

func TestBruteForceSingleByteXor(t *testing.T) {
	text := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := byte('X')
	xored := XorByte([]byte(text), key)

	decrypted, k := BruteForceSingleByteXor(xored)

	//fmt.Printf("key: '%c'\n", k)
	fmt.Println(string(decrypted))
	assert.AssertEq(key, k)
	assert.AssertEq(text, string(decrypted))

	fmt.Println("Ok")
}
