package hex

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func EncodeStr(s string) ([]byte, error) {
	var bytes []byte

	if len(s) <= 0 || len(s)%2 != 0 {
		msg := fmt.Sprintf("Invalid hex string: '%s'", s)
		return bytes, errors.New(msg)
	}

	for i := 0; i < len(s); i += 2 {
		tmp := string(s[i]) + string(s[i+1])
		i, err := strconv.ParseInt(tmp, 16, 8)
		check(err)

		bytes = append(bytes, byte(i))
	}

	return bytes, nil
}

func Decode(bytes []byte) string {
	var s string

	for _, b := range bytes {
		s += fmt.Sprintf("%x", b)
	}

	return s
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
