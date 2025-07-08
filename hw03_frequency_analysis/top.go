package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`([a-zA-Zа-яА-Я\-]+[-.,]{0,}[a-zA-Zа-яА-Я]{1,}|[a-zA-Zа-яА-Я]{1}|[\-]{2,})`)

type WordCount struct {
	Word  string
	Count int
}

func Top10(rs string) []string {
	if len(strings.TrimSpace(rs)) == 0 {
		return []string{}
	}

	words := re.FindAllString(rs, -1)

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
	for _, value := range words {
		result[strings.ToLower(value)]++
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
