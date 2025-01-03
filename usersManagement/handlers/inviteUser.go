package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"task.com/usersManagement/firebase"
	"task.com/usersManagement/models"
)

func InviteUser(w http.ResponseWriter, r *http.Request) {
	var invitation models.Invitation
	if err := json.NewDecoder(r.Body).Decode(&invitation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validar la fecha de expiración
	if invitation.ExpiresAt.Before(time.Now()) {
		http.Error(w, "La fecha de expiración debe ser en el futuro", http.StatusBadRequest)
		return
	}

	// Generar un nuevo ID para la invitación
	docRef := firebase.Client.Collection("invitation").NewDoc()
	invitation.ID = docRef.ID

	_, err := docRef.Set(context.Background(), invitation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invitation)
}
