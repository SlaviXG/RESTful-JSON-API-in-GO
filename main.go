package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// Running this example will spin up a server on port 8080,
	// and can be accessed at http://localhost:8080

	fmt.Println("Handling started...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := html.EscapeString(r.URL.Path)
		fmt.Fprintf(w, "Received request: %q", path)
		//For more information: fmt.Println(r.Header)
		fmt.Printf("Received request: %q\n", path)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
