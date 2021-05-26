package main

import (
	"fmt"
	"log"
	"time"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type MetaOrder struct {
	MetaOrderId        int                `json:"metaOrderId"`
	MetaOrderType      string             `json:"metaOrderType"`
	Status             string             `json:"status"`
	StopLossTakeProfit StopLossTakeProfit `json:"stopLossTakeProfit"`
	CreateDateTime     time.Time          `json:"createDateTime"`
	CreateUserName     string             `json:"createUserName"`
	LastUpdateDateTime time.Time          `json:"lastUpdateDateTime"`
	LastUpdateUserName string             `json:"lastUpdateUserName"`
}

type StopLossTakeProfit struct {
	StopLossPrice   float32 `json:"stopLossPrice"`
	TakeProfitPrice float32 `json:"takeProfitPrice"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func createMetaOrder(w http.ResponseWriter, r *http.Request) {
	var metaOrder MetaOrder

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&metaOrder); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
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
