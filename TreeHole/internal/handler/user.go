package handler

import (
	"TreeHole/internal/service"
	"encoding/json"
	"net/http"
)


type createUserInput struct {
	Email, Username string
}

func (h *handler) createUser(w http.ResponseWriter, r *http.Request) {
	var in createUserInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.CreateUser(r.Context(), in.Email, in.Username)
	if err == service.ErrInvalidEmail || err == service.ErrInvalidUsername {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err == service.ErrEmailTaken || err == service.ErrUsernameTaken {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
