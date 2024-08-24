package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"slices"

	"github.com/vilroi/gocryptopals/utils"
	"github.com/vilroi/gocryptopals/xor"
	"github.com/vilroi/goutils/assert"
)

const challenge string = "./6.txt"

func main() {
	data := base64Decode(challenge)
	potential_keysz := getPossibleKeySize(data)

	for _, keysz := range potential_keysz {
		transposed := transposeBlocks(data, keysz)

		var key []byte
		for _, block := range transposed {
			_, k := xor.BruteForceSingleByteXor(block)
			key = append(key, k)
		}

		if slices.Index(key, byte(0)) >= 0 {
			continue
		}

		fmt.Printf("key: '%s'\n\n", string(key))
		decrypted := xor.Xor(data, key)
		fmt.Println(string(decrypted))
	}

}

func transposeBlocks(data []byte, length int) [][]byte {
	var newSlice [][]byte

	for i := 0; i < len(data); i += length {
		newSlice = append(newSlice, data[i:i+length])
	}
	assert.AssertEq(length, len(newSlice[0]))

	transposed := make([][]byte, length)
	for _, entry := range newSlice {
		for i, c := range entry {
			transposed[i] = append(transposed[i], c)
		}
	}
	assert.AssertEq(length, len(transposed))

	return transposed
}

func getPossibleKeySize(data []byte) []int {
	type Pair struct {
		Distance int
		KeySize  int
	}

	/* calculate hamming distance for each potential key size */
	var keysize_vec []Pair
	for keysz := 2; keysz < 40; keysz++ {
		first := data[:keysz]
		second := data[keysz : keysz*2]

		distance := utils.CalcHamming(first, second) / keysz

		keysize_vec = append(keysize_vec, Pair{distance, keysz})
	}

	/* sort based on distance in ascending order*/
	slices.SortFunc(keysize_vec, func(a, b Pair) int {
		if a.Distance > b.Distance {
			return -1
		} else if a.Distance < b.Distance {
			return 1
		} else {
			return 0
		}
	})

	/* return top 5 */
	var potential_keysz []int
	for i := 0; i < 5; i++ {
		potential_keysz = append(potential_keysz, keysize_vec[i].KeySize)
	}

	return potential_keysz
}

func base64Decode(file string) []byte {
	data := readAll(file)
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(data)))

	n, err := base64.StdEncoding.Decode(decoded, data)
	check(err)

	return decoded[:n]
}

func readAll(file string) []byte {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	data, err := io.ReadAll(f)
	check(err)

	return data
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
