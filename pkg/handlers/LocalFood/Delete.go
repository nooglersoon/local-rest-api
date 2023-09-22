package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nooglersoon/locale-rest-api/pkg/mocks"
)

func DeleteLocalFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, localFood := range mocks.LocalFoods {
		if localFood.Id == id {
			mocks.LocalFoods = append(mocks.LocalFoods[:index], mocks.LocalFoods[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Successfully Deleted Singe Local Food")
			break
		}
	}
}
