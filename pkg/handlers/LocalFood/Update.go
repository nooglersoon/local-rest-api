package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nooglersoon/locale-rest-api/pkg/mocks"
	"github.com/nooglersoon/locale-rest-api/pkg/models"
)

func UpdateLocalFood(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedLocalFood models.LocalFood
	json.Unmarshal(body, &updatedLocalFood)

	for index, localFood := range mocks.LocalFoods {
		if localFood.Id == id {
			localFood.Name = updatedLocalFood.Name
			localFood.Address = updatedLocalFood.Address
			localFood.Latitude = updatedLocalFood.Latitude
			localFood.Longitude = updatedLocalFood.Longitude
			localFood.Description = updatedLocalFood.Description

			mocks.LocalFoods[index] = localFood
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode("Successfully Updated Singe Local Food")
			break
		}
	}
}
