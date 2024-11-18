package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Graph map[int]Vertex
type Vertex map[int]float64
type Towns map[int]string

var townsDir string
var graph Graph
var towns Towns

func init() {
	flag.StringVar(&townsDir, "dir", "public\\towns.txt", "a file containing the town directories")
	graph = make(Graph)
	towns = make(Towns)
}

func readTowns() {
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
	}
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
		dist, err := strconv.ParseFloat(split[2], 32)
		if err != nil {
			log.Fatal("invalid town index")
		}

		if _, ok := graph[from]; !ok {
			graph[from] = make(Vertex)
		}
		graph[from][to] = dist
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during reading distances: %v\n", err)
	}
}

func main() {
	flag.Parse()
	if townsDir == "" {
		log.Fatal("can't have empty towns dir")
	}

	readTowns()

}
