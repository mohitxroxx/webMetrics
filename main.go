package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	fmt.Println("Server is up and running at http://localhost:5000")
	log.Fatal(http.ListenAndServe("localhost:5000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Server is up and Running<h1>"))
}
