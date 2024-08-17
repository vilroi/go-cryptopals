package hex

import (
	"fmt"
	"testing"
)

func TestEncodeStr(t *testing.T) {
	s := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	bytes, err := EncodeStr(s)
	check(err)
	fmt.Println(bytes)
}

func TestDecodeStr(t *testing.T) {
	s := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	bytes, err := EncodeStr(s)
	check(err)

	res := Decode(bytes)
	if res != s {
		t.Fatalf("Expected: %s\nGot: %s\n", s, res)
	} else {
		t.Log("OK")
	}

}
