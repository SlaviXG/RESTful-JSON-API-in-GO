package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the basic endpoint!")
	fmt.Printf("The header of the basic endpoint request looks like:\n%s\n\n", r.Header)
}

func TodoIndex(w http.ResponseWriter, t *http.Request) {
	todos := Todos{
		Todo{Name: "Presentation"},
		Todo{Name: "Meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
