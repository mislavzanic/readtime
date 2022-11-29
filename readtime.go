package readtime

import (
	"math"
	"strings"
)

type ReadTime struct {
    WordCount uint64
	Seconds   uint64
}

func countWords(content string, options Options) float64 {
	var count float64 = 0.0
	words := strings.Split(content, " ")
	for _, word := range words {
		count += float64(len(word)) / float64(options.wordLength)
	}
	return float64(math.Round(count))
}

func CalcReadTime(content string, options Options) ReadTime {
	wordCount := countWords(content, options)
	seconds := (wordCount / float64(options.CalculateWPM())) * 60

	return ReadTime{
		WordCount: uint64(wordCount),
		Seconds: uint64(seconds),
	}
}
