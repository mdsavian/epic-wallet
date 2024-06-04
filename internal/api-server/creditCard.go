package apiserver

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mdsavian/budget-tracker-api/internal/types"
)

type CreateNewCreditCardInput struct {
	Name string `json:"name"`
}

func (s *APIServer) handleCreateCreditCard(w http.ResponseWriter, r *http.Request) {
	cardInput := CreateNewCreditCardInput{}

	if err := json.NewDecoder(r.Body).Decode(&cardInput); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	creditCard := &types.CreditCard{
		ID:        uuid.Must(uuid.NewV7()),
		Name:      cardInput.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := s.store.CreateCreditCard(creditCard); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, creditCard)
}

func (s *APIServer) handleGetCreditCard(w http.ResponseWriter, r *http.Request) {
	categories, err := s.store.GetCreditCard()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, categories)
}

func (s *APIServer) handleGetCreditCardByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	creditCard, err := s.store.GetCreditCardByName(name)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, creditCard)
}

func (s *APIServer) handleGetCreditCardById(w http.ResponseWriter, r *http.Request) {
	id, err := getAndParseIDFromRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	creditCard, err := s.store.GetCreditCardByID(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, creditCard)
}

func (s *APIServer) handleArchiveCreditCard(w http.ResponseWriter, r *http.Request) {
	id, err := getAndParseIDFromRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	if _, err := s.store.GetCreditCardByID(id); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.store.ArchiveCreditCard(id)
	log.Println(err)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "CreditCard archived successfully")
}