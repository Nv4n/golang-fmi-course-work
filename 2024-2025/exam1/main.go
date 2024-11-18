package main

import (
	"exam1/algorithm"
	"exam1/reader"
	"exam1/types"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type HtmlTown struct {
	Name string
	ID   int
}

type HtmlResult struct {
	FromID   int
	FromTown string
	ToID     int
	ToTown   string
	Dist     float64
}

var townsDir string
var graph *types.Graph
var towns types.Towns

func init() {
	flag.StringVar(&townsDir, "dir", "public\\towns.txt", "a file containing the town directories")
	towns = make(types.Towns)
}

func main() {
	flag.Parse()
	if townsDir == "" {
		log.Fatal("can't have empty towns dir")
	}

	nodes := reader.ReadFile(townsDir, towns)
	graph = types.NewGraph(nodes)

	types.InitializeDistances(*graph)

	graph = types.NewGraph(nodes)
	var townsHtmlData []HtmlTown
	for i, t := range towns {
		townsHtmlData = append(townsHtmlData, HtmlTown{Name: t, ID: i})
	}
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/page.go.html", "views/choose.go.html"))
		_ = tmpl.ExecuteTemplate(w, "Page", townsHtmlData)
	})

	http.HandleFunc("POST /dijkstra", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		from := r.PostForm.Get("from-town")
		to := r.PostForm.Get("to-town")
		if from == "" || to == "" {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		if from == to {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		//fmt.Println(from)
		//fmt.Printf(to)
		var source *types.Node
		var target *types.Node

		for i, node := range graph.Nodes {
			if id, err := strconv.Atoi(from); err == nil && node.ID == id {
				source = graph.Nodes[i]
				break
			}
		}
		for i, node := range graph.Nodes {
			if id, err := strconv.Atoi(to); err == nil && node.ID == id {
				target = graph.Nodes[i]
				break
			}
		}
		algorithm.Dijkstra(graph, source, target)
		fmt.Println()
		fmt.Printf("From ID: %v, Town: %v \n", source.ID, towns[source.ID])
		fmt.Printf("To ID: %v, Town: %v, Dist: %v \n", target.ID, towns[target.ID], target.Distance)

		res := HtmlResult{FromID: source.ID, FromTown: towns[source.ID], ToID: target.ID, ToTown: towns[target.ID],
			Dist: target.Distance}
		tmpl := template.Must(template.ParseFiles("views/page.go.html", "views/res.go.html"))
		_ = tmpl.ExecuteTemplate(w, "Page", res)
	})

	Example(nodes)
	fmt.Println("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func Example(nodes []*types.Node) {
	fmt.Println("THIS IS JUST EXAMPLE")
	graph = types.NewGraph(nodes)
	source := graph.Nodes[0]
	target := graph.Nodes[4]
	algorithm.Dijkstra(graph, source, target)

	fmt.Printf("From ID: %v, Town: %v \n", source.ID, towns[source.ID])
	fmt.Printf("To ID: %v, Town: %v, Dist: %v \n", target.ID, towns[target.ID], target.Distance)
	fmt.Println("END OF JUST EXAMPLE")
}
