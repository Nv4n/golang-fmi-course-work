package main

import (
	"fmt"
	"lab00/stringutils"
	"rsc.io/quote"
)

func main() {
	s := "Hello world"
	fmt.Println(s)
	goquote := quote.Go()
	fmt.Println(goquote)
	sreverse := stringutils.Reverse(s)
	fmt.Println(sreverse)
}
