package freq

import (
	"errors"
	"io"
	"log"
	"os"
	"unicode"
)

const sample_file = "../../data/sample.txt"

type AnalysisResult struct {
	CharCount  map[byte]int
	NoiseRatio float32
}

func CalcScore(data []byte) (float32, error) {
	var score float32

	charCount := CountChars(data)
	noiseRatio := calcNoiseRatio(charCount, len(data))

	if noiseRatio > 5.0 {
		return -1, errors.New("Text has too much noise")
	}

	freqTable := getFreqTable()
	for char, count := range charCount {
		score += freqTable[char] * float32(count)
	}

	return score, nil
}

func calcNoiseRatio(charCount map[byte]int, length int) float32 {
	var noiseCount int
	for char, _ := range charCount {
		if isalnum(char) || unicode.IsSpace(rune(char)) {
			continue
		}
		noiseCount += 1
	}

	return (float32(noiseCount) / float32(length)) * 100.0
}

func CountChars(data []byte) map[byte]int {
	charCount := make(map[byte]int)

	for _, c := range data {
		c = tolower(c)
		if count, ok := charCount[c]; ok {
			charCount[c] = count + 1
		} else {
			charCount[c] = 1
		}
	}

	return charCount
}

func isalnum(c byte) bool {
	if '0' <= c && c <= '9' {
		return true
	}

	if 'A' <= c && c <= 'Z' {
		return true
	}

	if 'a' <= c && c <= 'z' {
		return true
	}

	return false
}

func tolower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		c += 32
	}

	return c
}

func getFreqTable() map[byte]float32 {
	freq := make(map[byte]float32)
	data := readAll(sample_file)
	charCount := CountChars(data)

	for char, count := range charCount {
		freq[char] = float32(count) / float32(len(data))
	}

	return freq
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
