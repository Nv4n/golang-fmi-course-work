package main

import "flag"

type Graph map[int]Vertex
type Vertex map[int]int
type towns map[int]string

var townsDir string

func init() {
	flag.StringVar(&townsDir, "dir", "public/towns.txt", "a file containing the town directories")
}

func main() {

}
