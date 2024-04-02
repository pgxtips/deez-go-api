package models

import (
    "log"
    "math/rand"
    "deez-go-api/internal/db"
)


// public nut
type Nut struct {
    ID int `json:"id" db:"id"`
    Nut string `json:"nut" db:"joke"`
}

func GetAllNuts() ([]Nut, error){
    
    db := database.GetInstance()
    defer db.Close()

    query1 := "SELECT * FROM nuts.deez_nuts_jokes;"
    rows, err := db.Query(query1)

    if err != nil {
        log.Println(err)
        return nil, err
    }

    defer rows.Close()


    db_nuts := []Nut{}

    for rows.Next() {
        var n Nut 

        if err := rows.Scan(&n.ID, &n.Nut); err != nil {
            log.Println(err)
            return nil, err
        }

        db_nuts = append(db_nuts, n)
    }

    return db_nuts, nil
}

func GetNutByID(id int) (Nut, error) {

    db := database.GetInstance()
    defer db.Close()

    result_nut := Nut{}

    query1 := "SELECT * FROM nuts.deez_nuts_jokes WHERE id = $1 LIMIT 1;"
    err := db.Get(&result_nut, query1, id)

    if err != nil {
        log.Println(err)
        return Nut{}, err
    }

    return result_nut, nil
}

func GetNutCount() (int, error){

    db := database.GetInstance()
    defer db.Close()

    count := 0

    query1 := "SELECT COUNT(*) FROM nuts.deez_nuts_jokes;"
    err := db.Get(&count, query1)

    if err != nil {
        log.Println(err)
        return 0, err
    }

    return count, nil
}

func GetRandomNut() (Nut, error) {

    nut_count, err := GetNutCount()
    if err != nil { return Nut{}, err }

    nut_data, err := GetNutByID(rand.Intn(nut_count) + 1)
    if err != nil { return Nut{}, err }

    return nut_data, nil
}
