package controllers;

import (
    "fmt"
    "net/http"
    "encoding/json"
    "deez-go-api/internal/models"
)

func GetNuts(w http.ResponseWriter, r *http.Request) {
    random_number := 3;

    // get nut data where id is equal to random_number
    nut_data := models.GetNutByID(random_number)

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(nut_data) 

    if err != nil {
        fmt.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

}
