package freq

import (
	"errors"
	"io"
	"log"
	"os"
	"slices"
	"unicode"
)

const sample_file = "../../data/sample.txt"

type AnalysisResult struct {
	CharCount  []CharCount
	NoiseRatio float32
}

type CharCount struct {
	Char  byte
	Count int
}

func CalcScore(data []byte) (float32, error) {
	var score float32

	result := AnalyzeText(data)

	if result.NoiseRatio > 5.0 {
		return -1, errors.New("Text has too much noise")
	}

	freqTable := getFreqTable()
	for _, count := range result.CharCount {
		score += freqTable[count.Char] * float32(count.Count)
	}

	return score, nil
}

func AnalyzeText(data []byte) AnalysisResult {
	var result AnalysisResult

	result.CharCount = getCharCount(data)
	result.NoiseRatio = calcNoiseRatio(result.CharCount, len(data))

	return result
}

func calcNoiseRatio(charCount []CharCount, length int) float32 {
	var noiseCount int
	for _, entry := range charCount {
		if isalnum(entry.Char) || unicode.IsSpace(rune(entry.Char)) {
			continue
		}
		noiseCount += 1
	}

	return (float32(noiseCount) / float32(length)) * 100.0
}

func getCharCount(data []byte) []CharCount {
	charCount := CountChars(data)
	sorted := sortCharCount(charCount)

	return sorted
}

func sortCharCount(charCount map[byte]int) []CharCount {
	var countVec []CharCount

	for key, val := range charCount {
		countVec = append(countVec, CharCount{key, val})
	}

	slices.SortFunc(countVec, func(a, b CharCount) int {
		if a.Count > b.Count {
			return -1
		} else if a.Count < b.Count {
			return 1
		} else {
			return 0
		}
	})

	return countVec
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
