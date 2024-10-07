package web

import (
	"consumer-service/internal/usecase/find_by_id"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebAccountHandler struct {
	FindByIdAccountUseCase *find_by_id.FindByIdAccountUseCase
}

func NewWebAccountHandler(usecase find_by_id.FindByIdAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		FindByIdAccountUseCase: &usecase,
	}
}

func (h *WebAccountHandler) FindAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "account_id")
	output, err := h.FindByIdAccountUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
