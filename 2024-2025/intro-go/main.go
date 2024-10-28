package main

import (
	"fmt"
	w_c "intro-go/cmd/word_counting"
)

func main() {
	counts := make(map[string]int)
	var str string
	for {
		n, _ := fmt.Scanf("%s", &str)
		if n == 0 {
			break
		}

		w_c.WordCount(str, counts)
	}
	fmt.Printf("%v\n", counts)
}
