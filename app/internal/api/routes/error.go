package routes

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func RegisterErrorRoutes(router *http.ServeMux) {
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        err := json.NewEncoder(w).Encode("error") 

        if err != nil {
            fmt.Println(err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
    })
}

