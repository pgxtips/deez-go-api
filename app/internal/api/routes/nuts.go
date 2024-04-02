package routes

import (
    "log"
    "net/http"
    "deez-go-api/internal/api/controllers"
)


func RegisterNutRoutes(router *http.ServeMux) {
    log.Println("Registering Nut Routes")
    router.HandleFunc("/nuts", controllers.GetNutsHandler)
} 
