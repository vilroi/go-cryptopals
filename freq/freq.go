package freq

import (
	"errors"
	"slices"
	"strings"
	"unicode"
)

type CharCount struct {
	Char  string
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

	if calcNoiseRatio(charCount, len(data)) > 10.0 {
		return -1, errors.New("Text has too much noise")
	}

	compareLen := len(topChars) + 1
	for _, count := range charCount[:compareLen] {
		//if strings.IndexByte(topChars, count.Char) > -1 {
		if strings.Contains(topChars, count.Char) {
			score++
		}
	}

	compareLen = len(lastChars)
	for _, count := range charCount[len(charCount)-compareLen:] {
		//if strings.IndexByte(lastChars, count.Char) > -1 {
		if strings.Contains(lastChars, count.Char) {
			score++
		}
	}

	return score, nil
}

func calcNoiseRatio(charCount []CharCount, length int) float32 {
	var noiseCount int
	for _, entry := range charCount {
		if entry.Char == "noise" {
			noiseCount = entry.Count
			break
		}
	}

	return (float32(noiseCount) / float32(length)) * 100.0
}

func getCharCount(data []byte) []CharCount {
	charCount := countChars(data)
	sorted := sortCharCount(charCount)

	return sorted
}

func sortCharCount(charCount map[string]int) []CharCount {
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

func countChars(data []byte) map[string]int {
	charCount := make(map[string]int)
	for _, c := range data {
		if !unicode.IsPrint(rune(c)) {
			if count, ok := charCount["noise"]; ok {
				charCount["noise"] = count + 1
			} else {
				charCount["noise"] = 1
			}
			continue
		}
		if !unicode.IsLetter(rune(c)) {
			continue
		}

		//key := tolower(c)
		key := strings.ToLower(string(c))
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
