package controllers;

import (
    "fmt"
    "net/http"
    "encoding/json"
    "deez-go-api/internal/models"
)

func GetNutsHandler(w http.ResponseWriter, r *http.Request) {

    nut_data, nut_err := models.GetRandomNut()

    if nut_err != nil {
        fmt.Println(nut_err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(nut_data) 

    fmt.Println(nut_data)

    if err != nil {
        fmt.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}
