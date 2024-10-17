package main

import (
	"github.com/Masterminds/sprig/v3"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var t = template.Must(template.New("global").Funcs(sprig.FuncMap()).ParseFiles("templates/index.html", "templates/food.html"))
var reqChan chan struct{}

type Food struct {
	Id   int
	Name string
}

var Foods = []Food{
	{Id: 1, Name: "Ice Cream"},
	{Id: 2, Name: "Pizza"},
	{Id: 3, Name: "Chocolate"},
	{Id: 4, Name: "Cheseburger"},
	{Id: 5, Name: "Oreo"},
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("GET /else", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		rc := http.NewResponseController(w)

		if err := loadAsync(w, rc); err != nil {
			log.Printf("Error loading async content: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func loadAsync(w http.ResponseWriter, rc *http.ResponseController) (err error) {
	err = t.ExecuteTemplate(w, "template", nil)
	if err != nil {
		return
	}

	err = rc.Flush()
	if err != nil {
		return
	}

	err = t.ExecuteTemplate(w, "content", nil)
	if err != nil {
		return
	}

	err = rc.Flush()
	if err != nil {
		return
	}

	for i, item := range Foods {
		select {
		default:
			data := struct {
				Index int
				Value string
			}{i, item.Name}
			err = t.ExecuteTemplate(w, "slot", data)
			if err != nil {
				return
			}
			err = rc.Flush()
			if err != nil {
				return
			}

			time.Sleep(time.Second)
		}

	}

	err = t.ExecuteTemplate(w, "tail", nil)
	if err != nil {
		return
	}

	err = rc.Flush()
	if err != nil {
		return
	}
	return nil
}
