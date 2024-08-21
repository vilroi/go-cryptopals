package xor

import (
	"github.com/vilroi/gocryptopals/freq"
)

func XorStrings(text, key string) []byte {
	data := []byte(text)
	k := []byte(key)

	return Xor(data, k)
}

func XorByte(data []byte, key byte) []byte {
	var bytes []byte

	for _, c := range data {
		bytes = append(bytes, c^key)
	}

	return bytes
}

func Xor(data, key []byte) []byte {
	var bytes []byte
	i := 0
	for _, c := range data {
		bytes = append(bytes, c^key[i])
		i++

		if i == len(key) {
			i = 0
		}
	}

	return bytes
}

func BruteForceSingleByteXor(data []byte) ([]byte, byte) {
	var i byte
	var key byte
	var highScore int

	for ; i < 255; i++ {
		tmp := XorByte(data, i)
		score, err := freq.CalcScore(tmp)
		if err != nil {
			continue
		}

		if score > highScore {
			key = i
			highScore = score
		}
	}

	result := XorByte(data, key)

	return result, key
}
