package main

// To get the package: go get github.com/gorilla/mux
import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Running this example will spin up a server on port 8080,
	// and can be accessed at http://localhost:8080.
	// Run: go run github.com/SlaviXG/Practice-RESTful-JSON-API-in-GO.

	fmt.Println("Handling started...")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
