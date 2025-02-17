package api

import (
	"html/template"
	"net/http"
)

func Handler() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))
		data := "ToDo"
		temp.Execute(w, data)

	}

	http.HandleFunc("/", h1)

	http.ListenAndServe(":8000", nil)

}
