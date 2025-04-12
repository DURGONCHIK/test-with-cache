package controller

import (
	"encoding/json"
	"net/http"
	"testing/internal/usecase"
)

type EncryptHandler struct {
	useCase usecase.EncryptUseCase
}

func NewEncryptHandler(uc usecase.EncryptUseCase) *EncryptHandler {
	return &EncryptHandler{useCase: uc}
}

func (h *EncryptHandler) Encrypt(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Algorithm string `json:"algorithm"`
		Input     string `json:"input"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.useCase.Encrypt(request.Algorithm, request.Input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}
