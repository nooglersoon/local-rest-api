package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/nooglersoon/locale-rest-api/pkg/handlers/LocalFood"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello, Locale!")
	})

	router.HandleFunc("/localfoods", handlers.GetAllLocalFoods).Methods(http.MethodGet)
	router.HandleFunc("/localfood", handlers.AddLocalFood).Methods(http.MethodPost)
	router.HandleFunc("/localfoods/{id}", handlers.GetLocalFoodById).Methods(http.MethodGet)
	router.HandleFunc("/localfoods/{id}", handlers.UpdateLocalFood).Methods(http.MethodPut)
	router.HandleFunc("/localfoods/{id}", handlers.DeleteLocalFood).Methods(http.MethodDelete)

	log.Println("Locale Rest API is running!")
	http.ListenAndServe(":3333", router)
}
