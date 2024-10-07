package main

import (
	"fmt"
	"strings"
)

func count_words(s string) (result map[string]int) {
	result = make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		result[word]++
	}
	return
}

func main() {
	fmt.Println(count_words("Pesho,  otide za riba!!! Ne znaesh li tova.. znaesh ....."))
}
