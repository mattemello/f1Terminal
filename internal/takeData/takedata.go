package takedata

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const URLSite = "https://api.openf1.org/v1/"

var Now string
var Previus string

var Car CarData

func StartTicked() {
	now := time.Now().UTC()
	previus := now.Add(time.Duration(-1) * time.Second)

	Now = now.Format("2006-01-02 15:04:05")
	Previus = previus.Format("2006-01-02 15:04:05")

	Now = strings.ReplaceAll(Now, " ", "T")
	Previus = strings.ReplaceAll(Previus, " ", "T")

	fmt.Println("now: ", Now, " previus: ", Previus)
	GetDataCar()
}

// todo: change the error to only return
func GetDataCar() error {
	fmt.Println(URLSite + "car_data?date>" + Previus + "&date<=" + Now)
	obj, err := http.Get(URLSite + "car_data?date>" + Previus + "&date<=" + Now)
	if err != nil {
		log.Println("error in the get, ", err)
		return err
	}

	body, err := io.ReadAll(obj.Body)
	if err != nil {
		log.Println("error in the readAll: ", err)
		return err
	}

	err = json.Unmarshal(body, &Car)
	if err != nil {
		log.Println("error in the unmarshal: ", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		return err
	}

	fmt.Println(Car)

	return nil

}
