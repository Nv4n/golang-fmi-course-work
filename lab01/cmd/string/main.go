package main

import "fmt"

func main() {
	s := "abc[]码⼱"
	for index, rune := range s {
		fmt.Printf("%#U starting at %d\n", rune, index)
	}
	for index := 0; index < len(s); index++ {
		fmt.Printf("%#U starting at %d\n", s[index], index)
	}
	runes := []rune(s)
	runes[0] = 't'
	fmt.Println(string(runes))
}
