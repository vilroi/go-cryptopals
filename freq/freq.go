package freq

import (
	"errors"
	"slices"
	"strings"
	"unicode"
)

type CharCount struct {
	Char  byte
	Count int
}

const topChars string = "etaoinsrhld"
const lastChars string = "ybvkxjqz"

func CalcScore(data []byte) (int, error) {
	var score int

	charCount := getCharCount(data)
	minLen := len(topChars) + len(lastChars) + 2
	if len(charCount) < minLen {
		return -1, errors.New("Not enough valid characters in data")
	}

	compareLen := len(topChars) + 1
	for _, count := range charCount[:compareLen] {
		if strings.IndexByte(topChars, count.Char) > -1 {
			score++
		}
	}

	compareLen = len(lastChars)
	for _, count := range charCount[len(charCount)-compareLen:] {
		if strings.IndexByte(lastChars, count.Char) > -1 {
			score++
		}
	}

	return score, nil
}

func getCharCount(data []byte) []CharCount {
	charCount := countChars(data)
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

func countChars(data []byte) map[byte]int {
	charCount := make(map[byte]int)
	for _, c := range data {
		if !unicode.IsLetter(rune(c)) {
			continue
		}

		key := tolower(c)
		if count, ok := charCount[key]; ok {
			charCount[key] = count + 1
		} else {
			charCount[key] = 1
		}
	}

	return charCount
}

func tolower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		c += 32
	}

	return c
}
