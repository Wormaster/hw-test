package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
	"unicode"
)

var re = regexp.MustCompile(`(?i)(\p{L}+(?:\p{P}+\p{L}+)*|[\p{P}]{2,})`)

type WordCount struct {
	Word  string
	Count int
}

func Top10(rs string) []string {
	if len(strings.TrimSpace(rs)) == 0 {
		return []string{}
	}

	rs = strings.ReplaceAll(rs, "\n", " ")

	// words := re.FindAllString(rs, -1)
	words := strings.Split(rs, " ")

	wc := countWords(words)
	swc := sortMap(wc)

	length := len(swc)
	if length > 10 {
		length = 10
	}

	result := make([]string, length)
	for i, w := range swc {
		result[i] = w.Word
		if i == 9 {
			break
		}
	}

	return result
}

func countWords(words []string) map[string]int {
	result := make(map[string]int)
	var key string
	for _, value := range words {
		if len(EmojisOnly(value)) > 1 {
			key = value
		} else {
			key = re.FindString(value)
		}
		if len(strings.TrimSpace(key)) == 0 {
			continue
		}

		result[strings.ToLower(key)]++
	}

	return result
}

func sortMap(wc map[string]int) []WordCount {
	i := 0
	wcList := make([]WordCount, len(wc))
	for word, count := range wc {
		wcList[i] = WordCount{word, count}
		i++
	}

	sort.Slice(wcList, func(i, j int) bool {
		if wcList[i].Count != wcList[j].Count {
			return wcList[i].Count > wcList[j].Count
		}
		return wcList[i].Word < wcList[j].Word
	})

	return wcList
}

func isEmojiByExclusion(r rune) bool {
	return !unicode.IsLetter(r) &&
		!unicode.IsDigit(r) &&
		!unicode.IsPunct(r) &&
		!unicode.IsSpace(r)
}

func EmojisOnly(s string) string {
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if isEmojiByExclusion(r) {
			out = append(out, r)
		}
	}
	return string(out)
}
