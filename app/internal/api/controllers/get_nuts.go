package controllers;

import (
    "log"
    "net/http"
    "encoding/json"
    "deez-go-api/internal/models"
)

func GetNutsHandler(w http.ResponseWriter, r *http.Request) {

    nut_data, nut_err := models.GetRandomNut()

    if nut_err != nil {
        log.Println(nut_err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(nut_data) 

    log.Println(nut_data)

    if err != nil {
        log.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}
