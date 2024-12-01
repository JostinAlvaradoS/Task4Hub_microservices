package handlers

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "usersManagement/firebase"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    doc, err := firebase.Client.Collection("users").Doc(id).Get(context.Background())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(doc.Data())
}