package api 

import (
    "log"
    "net/http"
    "os"
    "deez-go-api/internal/api/routes"
)

func ServerStart() {
    router := http.NewServeMux()
    routes.RegisterNutRoutes(router)
    routes.RegisterErrorRoutes(router)

    port := os.Getenv("APP_PORT")

    log.Println("starting server on port "+port+"...")
    
    if err := http.ListenAndServe(":"+port, router); err != nil {
        log.Println(err)
    }
}
