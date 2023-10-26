package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
)

type Counts = map[string]int
type WordCount struct {
	Word  string
	Count int
}

var stopwordsList = []string{
	"a", "about", "above", "after", "again", "against", "all", "am", "an", "and", "any", "are",
	"aren't", "as", "at", "be", "because", "been", "before", "being", "below", "between", "both",
	"but", "by", "can't", "cannot", "could", "couldn't", "did", "didn't", "do", "does", "doesn't",
	"doing", "don't", "down", "during", "each", "few", "for", "from", "further", "had", "hadn't",
	"has", "hasn't", "have", "haven't", "having", "he", "he'd", "he'll", "he's", "her", "here",
	"here's", "hers", "herself", "him", "himself", "his", "how", "how's", "i", "i'd", "i'll", "i'm",
	"i've", "if", "in", "into", "is", "isn't", "it", "it's", "its", "itself", "let's", "me", "more",
	"most", "mustn't", "my", "myself", "no", "nor", "not", "of", "off", "on", "once", "only", "or",
	"other", "ought", "our", "ours", "ourselves", "out", "over", "own", "same", "shan't", "she",
	"she'd", "she'll", "she's", "should", "shouldn't", "so", "some", "such", "than", "that", "that's",
	"the", "their", "theirs", "them", "themselves", "then", "there", "there's", "these", "they",
	"they'd", "they'll", "they're", "they've", "this", "those", "through", "to", "too", "under",
	"until", "up", "very", "was", "wasn't", "we", "we'd", "we'll", "we're", "we've", "were", "weren't",
	"what", "what's", "when", "when's", "where", "where's", "which", "while", "who", "who's", "whom",
	"why", "why's", "with", "won't", "would", "wouldn't", "you", "you'd", "you'll", "you're", "you've",
	"your", "yours", "yourself", "yourselves",
}
var splitRegexp = regexp.MustCompile(`\W`)
var stopWordsMap = make(map[string]struct{}, len(stopwordsList))

func init() {
	for _, stopWord := range stopwordsList {
		stopWordsMap[stopWord] = struct{}{}
	}
}

func ExtractKeywords(file *os.File) (Counts, error) {
	scanner := bufio.NewScanner(file)
	counts := make(map[string]int)
	for scanner.Scan() {
		words := splitRegexp.Split(scanner.Text(), -1)

		for _, word := range words {
			if _, ok := stopWordsMap[word]; ok || len(word) <= 2 {
				continue
			}
			counts[word]++
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return counts, nil
}
func ProcessFile(fname string) (Counts, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ExtractKeywords(file)
}

func TopKeywords(counts Counts, topN int) []WordCount {
	wordCounts := make([]WordCount, len(counts))
	for word, count := range counts {
		wordCounts = append(wordCounts, WordCount{word, count})
	}
	slices.SortFunc(wordCounts, func(a, b WordCount) int {
		return a.Count - b.Count
	})
	return wordCounts[:topN]
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage keywords <filename>1...fname.2...")
	}

	files := os.Args[1:]
	for _, fname := range files {
		counts, err := ProcessFile(fname)
		if err != nil {
			fmt.Printf("error opening file %s: %v\n", fname, err)
			continue
		}
	}
	ExtractKeywords(file)
}
