package freq

import (
	"fmt"
	"io"
	"os"
	"testing"
)

/*
func TestCountChars(t *testing.T) {
	file := "alice.txt"
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	data, err := io.ReadAll(f)
	check(err)

	charCount := countChars(data)
	//pairs := mapToPairs(charCount)
	for c, count := range charCount {
		fmt.Printf("%c: %d\n", c, count)
	}
}
*/

func TestCalcScore(t *testing.T) {
	file := "alice.txt"
	data := readAll(file)

	charCount := getCharCount(data)
	for _, entry := range charCount {
		fmt.Printf("%c: %d\n", entry.Char, entry.Count)
	}
	fmt.Println(CalcScore(data))
}

/*
func TestToLower(t *testing.T) {
	for i := 0; i < 255; i++ {
		c := tolower(byte(i))
		fmt.Printf("%c\n", c)
	}
}

func TestCharFreq(t *testing.T) {
	f, err := os.Open("./t")
	check(err)
	defer f.Close()

	data, err := io.ReadAll(f)
	check(err)

	charFreq := calcCharFreq(data)
	for c, freq := range charFreq {
		fmt.Printf("%c: %f\n", c, freq)
	}
}

func TestBreakXor(t *testing.T) {
	f, err := os.Open("./t")
	check(err)
	defer f.Close()

	data, err := io.ReadAll(f)
	check(err)

	/*
		var key byte = 'y'
		data = xor.XorByte(data, key)

		var maxscore int
		var k byte
		for i := 1; i < 255; i++ {
			tmp := xor.XorByte(data, byte(i))
			score := CalcScore(tmp)

			fmt.Printf("'%c': %d\n", byte(i), score)
			if score > maxscore {
				maxscore = score
				k = byte(i)
			}
		}
		fmt.Printf("'%c': %d\n", k, maxscore)
}
*/

func readAll(path string) []byte {
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	data, err := io.ReadAll(f)
	check(err)

	return data
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
