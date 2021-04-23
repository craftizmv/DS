package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//mux router - Handles all the incoming requests and redirects them.
	router := mux.NewRouter()

	//Handler
	router.HandleFunc("/test", test)

	//Start the web server
	http.ListenAndServe(":5000", router)

}

func test(w http.ResponseWriter, r *http.Request) {
	//Sending a slice of byte
	// w.Write([]byte("This is a test response"))

	w.Header().Set("Content-Type", "application/json")
	//sending the json response using the json encoder.
	//Also we are sending the response by encoding a anonymous struct
	json.NewEncoder(w).Encode(struct {
		ID string
	}{"111"})

}
