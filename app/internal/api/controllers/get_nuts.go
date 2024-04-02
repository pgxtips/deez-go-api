package controllers;

import (
    "fmt"
    "net/http"
    "encoding/json"
    "math/rand"
    "deez-go-api/internal/models"
)

func GetNuts(w http.ResponseWriter, r *http.Request) {

    // random number between 1 - 62
    random_number := rand.Intn(62) + 1

    // get nut data where id is equal to random_number
    nut_data := models.GetNutByID(random_number)

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(nut_data) 

    fmt.Println(nut_data)

    if err != nil {
        fmt.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

}
