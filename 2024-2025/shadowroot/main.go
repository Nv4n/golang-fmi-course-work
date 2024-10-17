package shadowroot

import (
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var t = template.Must(template.New("global").Funcs(sprig.FuncMap()).ParseFiles("templates/index.html", "templates/food.html"))

type Food struct {
	Id   int
	Name string
}

func sendFoodInDelayedOrder(data []Food, order []int, ch chan Food) {
	go func() {
		for _, index := range order {
			time.Sleep(5 * time.Second)
			ch <- data[index]
		}
		close(ch)
	}()
}

var Foods = []Food{
	{Id: 1, Name: "Ice Cream"},
	{Id: 2, Name: "Pizza"},
	{Id: 3, Name: "Chocolate"},
	{Id: 4, Name: "Cheseburger"},
	{Id: 5, Name: "Oreo"},
}

type AwaitedSlot struct {
	Slot string
	Html string
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

		err := t.ExecuteTemplate(w, "template", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Flush the writer to send the initial content
		err = rc.Flush()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.ExecuteTemplate(w, "content", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Flush the writer to send the initial content
		err = rc.Flush()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Simulate slow-loading content
		time.Sleep(2 * time.Second)

		// Stream additional content
		for i, item := range []string{"Item 1", "Item 2", "Item 3"} {
			//t.ExecuteTemplate(w, "slot")
			_, err = fmt.Fprintf(w, "<p slot=\"item-%v\">%s</p>", i, item)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = rc.Flush()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			time.Sleep(time.Second)
		}

		err = t.ExecuteTemplate(w, "tail", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Flush the writer to send the initial content
		err = rc.Flush()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
