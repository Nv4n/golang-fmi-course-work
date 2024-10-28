package word_counting

import (
	"fmt"
	"strings"
)

func WordCount(s string, wc map[string]int) map[string]int {

	for _, word := range strings.Fields(s) {
		fmt.Println(word)
		wc[word]++
	}

	//for word, count := range wc {
	//
	//}
	//sort.Slice()
	return wc
}
