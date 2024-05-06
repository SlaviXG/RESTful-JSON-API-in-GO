package main

// To get the package: go get github.com/gorilla/mux
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Running this example will spin up a server on port 8080,
	// and can be accessed at http://localhost:8080

	fmt.Println("Handling started...")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the basic endpoint!")
	fmt.Printf("The header of the request looks like:\n%s\n\n", r.Header)
}

func TodoIndex(w http.ResponseWriter, t *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
