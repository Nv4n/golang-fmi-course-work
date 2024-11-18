package main

import (
	"bufio"
	"exam1/types"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Towns map[int]string

var townsDir string
var graph types.Graph
var towns Towns

func init() {
	flag.StringVar(&townsDir, "dir", "public\\towns.txt", "a file containing the town directories")
	towns = make(Towns)
}

func readFile() {
	pwd, _ := os.Getwd()
	open, err := os.Open(fmt.Sprintf("%s\\%s", pwd, townsDir))
	if err != nil {
		log.Fatalf("error opening fileDir %s: %v", townsDir, err)
	}
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(open)
	scanner := bufio.NewScanner(open)

	var nodes []*types.Node

	nodes = readTowns(scanner, nodes)
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during reading towns: %v\n", err)
	}

	for scanner.Scan() {
		distance := scanner.Text()
		split := strings.Split(distance, ",")
		if len(split) != 3 {
			log.Fatal("invalid distance format")
		}
		from, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal("invalid town index")
		}
		to, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("invalid town index")
		}
		dist, err := strconv.ParseFloat(split[2], 64)
		if err != nil {
			log.Fatal("invalid town index")
		}
		var fromInd int
		var toInd int
		for i, n := range nodes {
			if n.ID == from {
				fromInd = i
				break
			}
		}
		for i, n := range nodes {
			if n.ID == to {
				toInd = i
				break
			}
		}

		nodes[fromInd].Neighbors = append(
			nodes[fromInd].Neighbors,
			&types.Edge{Destination: nodes[toInd], Weight: dist})
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during reading distances: %v\n", err)
	}
}

func readTowns(scanner *bufio.Scanner, nodes []*types.Node) []*types.Node {
	for scanner.Scan() {
		town := scanner.Text()
		if town == "" || town == "\n" {
			break
		}
		split := strings.Split(town, ",")
		if len(split) != 2 {
			log.Fatal("invalid town format")
		}
		ind, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal("invalid town index")
		}
		towns[ind] = split[1]

		node := &types.Node{
			ID:       ind,
			Distance: math.Inf(1),
			Visited:  false,
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func main() {
	flag.Parse()
	if townsDir == "" {
		log.Fatal("can't have empty towns dir")
	}

	readFile()

}
