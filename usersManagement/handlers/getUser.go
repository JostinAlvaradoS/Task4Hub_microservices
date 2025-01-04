package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"task.com/usersManagement/firebase"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uid := params["uid"]
	// Realizar una consulta en Firestore para buscar por el campo "uid"
	iter := firebase.Client.Collection("user").Where("UID", "==", uid).Documents(context.Background()) //cambio gcloud
	defer iter.Stop()

	// Obtener el primer documento que coincida con el uid
	doc, err := iter.Next()
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	// Codificar los datos del usuario encontrados
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doc.Data())
}
