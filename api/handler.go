package api

import (
	"fmt"
	"net/http"
)

func Handler() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")

	}

	http.HandleFunc("/", h1)

	http.ListenAndServe(":8000", nil)

}
