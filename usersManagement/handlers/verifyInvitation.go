package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"task.com/usersManagement/firebase"
	"task.com/usersManagement/models"
)

func VerifyInvitation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	invitationID := params["id"]

	// Buscar el documento por el ID
	docRef := firebase.Client.Collection("invitation").Doc(invitationID)
	doc, err := docRef.Get(context.Background())
	if err != nil {
		http.Error(w, "Invitación no encontrada", http.StatusNotFound)
		return
	}

	var invitation models.Invitation
	if err := doc.DataTo(&invitation); err != nil {
		http.Error(w, "Error al decodificar la invitación", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invitation)
}
