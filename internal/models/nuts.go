package models

type Nut struct {
    ID int `json:"id"`
    Nut string `json:"nut"`
}

func GetAllNuts() []Nut{
    return []Nut{
        {ID: 0, Nut: "deez nuts get em 0"},
        {ID: 1, Nut: "deez nuts get em 1"},
        {ID: 2, Nut: "deez nuts get em 2"},
        {ID: 3, Nut: "deez nuts get em 3"},
        {ID: 4, Nut: "deez nuts get em 4"},
        {ID: 5, Nut: "deez nuts get em 5"},
        {ID: 6, Nut: "deez nuts get em 6"},
    }
}

func GetNutByID(id int) Nut {
    all_nuts := GetAllNuts()
    for _, nut := range all_nuts {
        if nut.ID == id {
            return nut
        }
    }
    return Nut{}
}
