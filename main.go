package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type TodoPage struct {
	PageTitle string
	Task      string
}

// type TodoTask struct {
// }

func main() {
	fmt.Println("ToDo")
	m := make(map[int]string)
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		task := r.PostFormValue("task")
		m[0] = task
		title := "Todo"
		tmpl.Execute(w, TodoPage{title, task})
	})
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":8080", nil)

}
