package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createMetaOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create Order")
}

func getMetaOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Read Order")
}

func deleteMetaOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Order")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/create/{msg}", createMetaOrder).Methods("POST")
	myRouter.HandleFunc("/{msg}", getMetaOrder).Methods("GET")
	myRouter.HandleFunc("/delete/{msg}", deleteMetaOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

func main() {
	fmt.Println("Listening on port 8080")
	handleRequests()
}
