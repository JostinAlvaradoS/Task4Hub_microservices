package handlers

import (
    "context"
    "net/http"

    "github.com/gorilla/mux"
   "usersManagement/firebase"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    _, err := firebase.Client.Collection("users").Doc(id).Delete(context.Background())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}