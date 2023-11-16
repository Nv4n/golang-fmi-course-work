package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	graph "travel_graph/structures"
)

func readNames(file *os.File) map[int]string {

	townNames := make(map[int]string)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text := reader.Text()
		if len(text) == 0 {
			return townNames
		}
		keyVal := strings.Split(text, ",")
		index, err := strconv.Atoi(keyVal[0])
		if err != nil {
			return nil
		}
		townNames[index] = keyVal[1]
	}
	return townNames

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Insufficient args count")
		return
	}
	filename := os.Args[1]
	pwd, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%v\\cmd\\%v", pwd, filename))
	if err != nil {
		return
	}

	test := graph.NewGraph()
	test.AddVertex(1)
	test.AddVertex(2)
	test.AddEdge(1, 2, 34)
	test.FindShortestPath(1, 2)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
}
