package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/usersManagement/firebase"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	firebase.InitFirebase()
	params := mux.Vars(r)
	uid := params["uid"]
	// Realizar una consulta en Firestore para buscar por el campo "uid"
	iter := firebase.Client.Collection("user").Where("uid", "==", uid).Documents(context.Background())
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
