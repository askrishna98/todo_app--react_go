package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aswin/go-react-todo/router"
)

// Entry point of the application
// initializes the router and starts the HTTP server on port 9000

func main() {
	r := router.Router()
	fmt.Println("starting the server on port 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
