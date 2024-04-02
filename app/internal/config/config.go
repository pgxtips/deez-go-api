package config

import(
    "os"
    "log"
    "github.com/joho/godotenv"
)


func ConfigInit(){

    err := godotenv.Load()

    if err != nil {
        log.Println("Error loading .env file")
    }

    port := os.Getenv("APP_PORT")
    if port == "" { log.Fatal("APP_PORT is not set") }

    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" { log.Fatal("DB_HOST is not set") }

    dbPort := os.Getenv("DB_PORT")
    if dbPort == "" { log.Fatal("DB_PORT is not set") }

    dbName := os.Getenv("DB_NAME")
    if dbName == "" { log.Fatal("DB_NAME is not set") }

    dbUser := os.Getenv("DB_USER")
    if dbUser == "" { log.Fatal("DB_USER is not set") }

    dbPassword := os.Getenv("DB_PASSWORD")
    if dbPassword == "" { log.Fatal("DB_PASSWORD is not set") }

    dbSslmode := os.Getenv("DB_SSLMODE")
    if dbSslmode == "" { log.Fatal("DB_SSLMODE is not set") }

    os.Setenv("APP_PORT", port)
    os.Setenv("DB_HOST", dbHost)
    os.Setenv("DB_PORT", dbPort)
    os.Setenv("DB_NAME", dbName)
    os.Setenv("DB_USER", dbUser)
    os.Setenv("DB_PASSWORD", dbPassword)
    os.Setenv("DB_SSLMODE", dbSslmode)

}


