package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	graph "travel_graph/structures"
)

func readFile(file *os.File) struct {
	Towns map[int]string
	Paths []struct {
		From, To int
		Distance float64
	}
} {

	townNames := make(map[int]string)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text := reader.Text()
		if len(text) == 0 {
			break
		}
		keyVal := strings.Split(text, ",")
		index, err := strconv.Atoi(keyVal[0])
		if err != nil {
			log.Fatal(err)
		}
		townNames[index] = keyVal[1]
	}
	var paths []struct {
		From, To int
		Distance float64
	}
	for reader.Scan() {
		text := reader.Text()

		fromToDist := strings.Split(text, ",")
		from, err := strconv.Atoi(fromToDist[0])
		if err != nil {
			log.Fatal(err)
		}
		to, err := strconv.Atoi(fromToDist[1])
		if err != nil {
			log.Fatal(err)
		}
		dist, err := strconv.ParseFloat(fromToDist[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		paths = append(paths, struct {
			From, To int
			Distance float64
		}{from, to, dist})
	}

	return struct {
		Towns map[int]string
		Paths []struct {
			From, To int
			Distance float64
		}
	}{townNames, paths}
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

	input := readFile(file)
	towns := input.Towns
	paths := input.Paths
	g := graph.NewGraph()
	for i, _ := range towns {
		g.AddVertex(i)
	}
	for _, path := range paths {
		g.AddEdge(path.From, path.To, path.Distance)
	}
	fmt.Println(g)
	g.FindShortestPath(2, 4)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
}
