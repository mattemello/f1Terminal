package takedata

import (
	"encoding/json"
	"fmt"
	"log"
)

// NOTE: this is a work in progress, idk if i will use this

func CarFunc() string {
	var car Car
	car.URL = URLSite + "car_data?date>" + Previus + "&date<=" + Now

	body, err := getData(car.URL)
	if err != nil {
		log.Println("error in the get, ", err)
		return ""
	}

	err = json.Unmarshal(body, &car.CarData)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Println(string(body))
	}

	fmt.Println(car.CarData)
	//note: here I update the tui
	return carToString(car.CarData)
}

func carToString(car CarData) string {
	return fmt.Sprintf("Gear: %d\nDrs: %d\nBrake: %d\nSpeed: %d\n", car[0].NGear, car[0].Drs, car[0].Brake, car[0].Speed)
}
