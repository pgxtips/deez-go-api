package database

import (
    "fmt"
    "log"
    "os"
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)


func connect() *sqlx.DB {

    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")
    user :=  os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    sslmode := os.Getenv("DB_SSLMODE")

    conn_str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        host, port, user, password, dbname, sslmode,
    )

    db, err := sqlx.Connect("postgres", conn_str)
    if err != nil {
        log.Fatalln(err)
    }

    return db
}

func GetInstance() *sqlx.DB {
    return connect()
}
