package takedata

import (
	"encoding/json"
	"errors"
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

func TakeCircuit() (Circuit, error) {
	var cir []Circuit
	circuitUrl := URLSite + "meetings?meeting_key=latest"

	body, err := GetData(circuitUrl)
	if err != nil {
		return Circuit{}, err
	}

	err = json.Unmarshal(body, &cir)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			return Circuit{}, errors.New(fmt.Sprintf("syntax error at byte offset %d \n body: %s", e.Offset, string(body)))
		}
		return Circuit{}, err
	}

	return cir[0], nil
}

// NOTE: the int in the map is the number of the Driver
func Interval() map[int]IntervalAll {
	var inter []IntervalAll
	intUrl := URLSite + "intervals?session_key=latest&date>" + Previus + "&date<=" + Now

	body, err := GetData(intUrl)
	if err != nil {
		log.Println("error in the get, ", err)
		return nil
	}

	err = json.Unmarshal(body, &inter)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Println(string(body))
		return nil
	}

	return cleanInterval(inter)
}

func cleanInterval(interval []IntervalAll) map[int]IntervalAll {
	var intervalMap = make(map[int]IntervalAll)

	for _, elem := range interval {
		value, in := intervalMap[elem.DriverNumber]
		if !in {
			intervalMap[elem.DriverNumber] = elem
			continue
		}

		if !elem.Date.After(value.Date) {
			intervalMap[elem.DriverNumber] = elem
			continue
		}
	}

	return intervalMap
}

func TickedDone() [][]string {
	now := time.Now().UTC()
	previus := now.Add(time.Duration(-1) * time.Second)

	Now = strings.ReplaceAll(now.Format("2006-01-02 15:04:05"), " ", "T")
	Previus = strings.ReplaceAll(previus.Format("2006-01-02 15:04:05"), " ", "T")

	session := Session()
	inter := Interval()

	return changedTable(session, inter)
}

func changedTable(clSe map[int]Position, inte map[int]IntervalAll) [][]string {
	var driv = make([][]string, 20)

	for _, elem := range clSe {
		driv[elem.Position-1] = []string{
			fmt.Sprintf("%d", elem.Position),
			drvMap[elem.DriverNumber].FirstName,
			drvMap[elem.DriverNumber].LastName,
			fmt.Sprintf("%d", elem.DriverNumber),
			fmt.Sprintf("%f", inte[elem.DriverNumber].GapToLeader),
			fmt.Sprintf("%f", inte[elem.DriverNumber].Interval),
			drvMap[elem.DriverNumber].TeamName,
		}
	}

	return nil
}

func CarFunc() string {
	var car Car
	car.URL = URLSite + "car_data?date>" + Previus + "&date<=" + Now

	body, err := GetData(car.URL)
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

func GetData(url string) ([]byte, error) {
	obj, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer obj.Body.Close()

	body, err := io.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
