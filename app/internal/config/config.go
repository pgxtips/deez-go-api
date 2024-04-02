package config

import(
    "os"
    "log"
    "github.com/joho/godotenv"
)


func ConfigInit(){

    err := godotenv.Load()

    if err != nil {
        log.Fatal("Error loading .env file")
    }

    os.Setenv("DB_HOST", os.Getenv("DB_HOST"))
    os.Setenv("DB_PORT", os.Getenv("DB_PORT"))
    os.Setenv("DB_NAME", os.Getenv("DB_NAME"))
    os.Setenv("DB_USER", os.Getenv("DB_USER"))
    os.Setenv("DB_PASSWORD", os.Getenv("DB_PASSWORD"))
    os.Setenv("DB_SSLMODE", "require")
}


