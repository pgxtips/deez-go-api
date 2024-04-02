package tests

import (
    "testing"
    "deez-go-api/internal/models"
    "deez-go-api/internal/config"
)

func Init(){
    config.ConfigInit()
}

func TestGetNutRandomNut(t *testing.T){
    Init()

    nut_data1, err1 := models.GetRandomNut()
    nut_data2, err2 := models.GetRandomNut()
    nut_data3, err3 := models.GetRandomNut()
    nut_data4, err4 := models.GetRandomNut()
    nut_data5, err5 := models.GetRandomNut()

    nut_data_arr := []models.Nut{nut_data1, nut_data2, nut_data3, nut_data4, nut_data5}
    nut_err_arr := []error{err1, err2, err3, err4, err5}

    for _, err := range nut_err_arr {
        if err != nil {
            t.Errorf("Error: %v", err)
        }
    }

    for _, nut := range nut_data_arr {
        if nut.Nut == "" {
            t.Errorf("Nut is empty")
        }
    }

    if nut_data1.Nut == nut_data2.Nut && nut_data2.Nut == nut_data3.Nut && nut_data3.Nut == nut_data4.Nut && nut_data4.Nut == nut_data5.Nut {
        t.Errorf("All nuts are the same")
    }

    t.Logf("Successfully got 5 random nuts")
}


func TestGetNutByID(t *testing.T){
    Init()

    expected_nut := []string{
        "What’s a squirrel’s favorite way to get injured? By busting deez nuts!",
        "What do you call nuts on a wall? Walnuts. What do you call nuts on your chest? Chestnuts. What do you call nuts on your chin? Deez nuts!",
        "Why didn’t the nut go to Dunkin Donuts? He’d rather be Dunkin Deez Nuts!",
        "Why did the nut go to therapy? To deal with deez nutty issues!",
        "What’s a nut’s favorite music genre? Hip-hop… they don’t drop the bass, they love dropping deez nuts!",
    }

    nut_data1, err1 := models.GetNutByID(1)
    nut_data2, err2 := models.GetNutByID(2)
    nut_data3, err3 := models.GetNutByID(3)
    nut_data4, err4 := models.GetNutByID(4)
    nut_data5, err5 := models.GetNutByID(5)

    nut_data_arr := []models.Nut{nut_data1, nut_data2, nut_data3, nut_data4, nut_data5}
    nut_err_arr := []error{err1, err2, err3, err4, err5}

    for _, err := range nut_err_arr {
        if err != nil {
            t.Errorf("Error: %v", err)
        }
    }

    for idx, nut := range nut_data_arr {
        if nut.Nut != expected_nut[idx] {
            t.Errorf("Expected nut: %v, Got: %v", expected_nut[idx], nut.Nut)
        }        
    }

    t.Logf("Successfully got 5 random nuts")
}
