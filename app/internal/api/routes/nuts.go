package routes

import (
    "fmt"
    "net/http"
    "deez-go-api/internal/api/controllers"
)


func RegisterNutRoutes(router *http.ServeMux) {
    fmt.Println("registering nut routes...")
    router.HandleFunc("/nuts", controllers.GetNuts)
} 
