package api 

import (
    "log"
    "net/http"
    "deez-go-api/internal/api/routes"
)

func ServerStart() {
    log.Println("starting server on port 8080...")

    router := http.NewServeMux()
    routes.RegisterNutRoutes(router)
    routes.RegisterErrorRoutes(router)

    
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Println(err)
    }
}
