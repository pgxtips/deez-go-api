package models

import (
    "fmt"
    "deez-go-api/internal/db"
)


// public nut
type Nut struct {
    ID int `json:"id" db:"id"`
    Nut string `json:"nut" db:"joke"`
}

func GetAllNuts() []Nut{
    
    db := database.GetInstance()
    defer db.Close()

    query1 := "SELECT * FROM nuts.deez_nuts_jokes;"
    rows, err := db.Query(query1)

    if err != nil {
        fmt.Println(err)
    }

    defer rows.Close()


    db_nuts := []Nut{}

    for rows.Next() {
        var n Nut 

        if err := rows.Scan(&n.ID, &n.Nut); err != nil {
            fmt.Println(err)
        }

        db_nuts = append(db_nuts, n)
    }

    return db_nuts
}

func GetNutByID(id int) Nut {

    db := database.GetInstance()
    defer db.Close()

    result_nut := Nut{}

    query1 := "SELECT * FROM nuts.deez_nuts_jokes WHERE id = $1 LIMIT 1;"
    err := db.Get(&result_nut, query1, id)

    if err != nil {
        fmt.Println(err)
        return Nut{}
    }

    return result_nut
}

func GetNutCount() int{

    db := database.GetInstance()
    defer db.Close()

    count := 0

    query1 := "SELECT COUNT(*) FROM nuts.deez_nuts_jokes;"
    err := db.Get(&count, query1)

    if err != nil {
        fmt.Println(err)
        return 0
    }

    return count 
}
