package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kashyapt/MongoDB/router"
)

func main() {
	r := router.Router()                        //initialize the router
	fmt.Println("Server is getting started...") // This line prints "Server is getting started..." to the
	//standard output,indicating that the HTTP server is starting up.
	fmt.Println("Listing at port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))
	// This line starts an HTTP server listening on port 4000 using the http.ListenAndServe() function.
	// r is typically an instance of a http.Handler or http.HandlerFunc that specifies how to handle incoming
	// HTTP requests. It could be a router or a handler function.
}
