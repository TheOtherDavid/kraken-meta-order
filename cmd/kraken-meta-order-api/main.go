// Version 0.2.0
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

		metaOrderId := mux.Vars(r)["id"]

		returnedMetaOrder, err := db.GetMetaOrder(metaOrderId)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error reading from DB")
			return
		}

		defer r.Body.Close()

		json.NewEncoder(w).Encode(returnedMetaOrder)
	}
}

func findMetaOrders() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		status := r.URL.Query().Get("status")

		searchCriteria := models.SearchCriteria{
			Status: status,
		}

		returnedMetaOrders, err := db.FindMetaOrders(searchCriteria)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error reading from DB")
			return
		}

		defer r.Body.Close()

		json.NewEncoder(w).Encode(returnedMetaOrders)
	}
}

func deleteMetaOrder() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		metaOrderId := mux.Vars(r)["id"]

		returnedMetaOrders, err := db.DeleteMetaOrder(metaOrderId)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error reading from DB")
			return
		}

		defer r.Body.Close()

		json.NewEncoder(w).Encode(returnedMetaOrders)
	}
}

type healthCheckResponse struct {
	Status string `json:"status"`
}

func health() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response := healthCheckResponse{
			Status: "Ok",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/{id}", deleteMetaOrder()).Methods("DELETE")
	myRouter.HandleFunc("/find", findMetaOrders()).Methods("GET")
	myRouter.HandleFunc("/health", health()).Methods("GET")
	myRouter.HandleFunc("/", createMetaOrder()).Methods("POST")
	myRouter.HandleFunc("/{id}", getMetaOrder()).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Listening on port 8080")
	handleRequests()
}
