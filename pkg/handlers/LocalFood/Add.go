package handlers

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/nooglersoon/locale-rest-api/pkg/mocks"
	"github.com/nooglersoon/locale-rest-api/pkg/models"
)

func AddLocalFood(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var localFood models.LocalFood
	json.Unmarshal(body, &localFood)

	// Append to the Book mocks
	localFood.Id = rand.Intn(100)
	mocks.LocalFoods = append(mocks.LocalFoods, localFood)

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Successfully Created Singe Local Food")
}
