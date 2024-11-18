package reader

import (
	"bufio"
	"exam1/types"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(townsDir string, towns types.Towns) []*types.Node {
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

	nodes = readTowns(scanner, nodes, towns)
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during reading towns: %v\n", err)
	}

	readDistances(scanner, nodes)
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during reading distances: %v\n", err)
	}
	return nodes
}

func readDistances(scanner *bufio.Scanner, nodes []*types.Node) {
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
}

func readTowns(scanner *bufio.Scanner, nodes []*types.Node, towns types.Towns) []*types.Node {
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
			ID: ind,
		}
		nodes = append(nodes, node)
	}
	return nodes
}
