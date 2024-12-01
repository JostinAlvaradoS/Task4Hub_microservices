package handlers

import (
    "context"
    "encoding/json"
    "net/http"

    "usersManagement/firebase"
    "usersManagement/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, _, err := firebase.Client.Collection("users").Add(context.Background(), user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}