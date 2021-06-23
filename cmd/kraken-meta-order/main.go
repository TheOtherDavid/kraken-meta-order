package main

import (
	"fmt"
	"log"

	"encoding/json"

	"net/http"

	"github.com/TheOtherDavid/kraken-meta-order/internal/db"
	"github.com/TheOtherDavid/kraken-meta-order/internal/models"

	"github.com/gorilla/mux"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func createMetaOrder() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var metaOrder models.MetaOrder

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&metaOrder); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		returnedMetaOrder, err := db.CreateMetaOrder(metaOrder)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error saving to DB")
			return
		}

		defer r.Body.Close()

		json.NewEncoder(w).Encode(returnedMetaOrder)
	}
}

func getMetaOrder() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		metaOrderId := mux.Vars(r)["msg"]

		returnedMetaOrder, err := db.GetMetaOrder(metaOrderId)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error reading from DB")
			return
		}

		defer r.Body.Close()

		json.NewEncoder(w).Encode(returnedMetaOrder)
	}
}

func deleteMetaOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Order")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", createMetaOrder()).Methods("POST")
	myRouter.HandleFunc("/{msg}", getMetaOrder()).Methods("GET")
	myRouter.HandleFunc("/delete/{msg}", deleteMetaOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

func main() {
	fmt.Println("Listening on port 8080")
	handleRequests()
}
