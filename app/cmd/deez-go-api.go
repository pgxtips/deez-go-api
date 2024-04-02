package main

//import "fmt"
import (
    "deez-go-api/internal/api"
    "deez-go-api/internal/config"
)


func main() {
    config.ConfigInit()
    api.ServerStart()
}
