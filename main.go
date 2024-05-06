package main

// To get the package: go get github.com/gorilla/mux
import (
	"fmt"
	"html"
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
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	path := html.EscapeString(r.URL.Path)
	fmt.Fprintf(w, "Received request: %q", path)
	//For more information: fmt.Println(r.Header)
	fmt.Printf("Received request: %q\n", path)
}
