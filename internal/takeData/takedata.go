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

func TakeDriverInSession() map[int]Driver {
	var driver []DriverAll
	drivUrl := URLSite + "drivers?session_key=latest"

	body, err := GetData(drivUrl)
	if err != nil {
		log.Println("error in the get, ", err)
		return nil
	}

	err = json.Unmarshal(body, &driver)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Println("error in the unmarshal: ", err, " \nbody: ", string(body))
	}

	return driverMap(driver)
}

var drvMap = make(map[int]Driver)

func driverMap(dr []DriverAll) map[int]Driver {

	for _, elem := range dr {
		drvMap[elem.DriverNumber] = Driver{
			FirstName:    elem.FirstName,
			LastName:     elem.LastName,
			NameAcronym:  elem.NameAcronym,
			DriverNumber: elem.DriverNumber,
			TeamName:     elem.TeamName,
		}
	}

	return drvMap
}

func IsSessionOn() bool {
	var session []Session
	sessionURL := URLSite + "sessions?session_key=latest&meeting_key=latest"

	body, err := GetData(sessionURL)

	if err != nil {
		log.Println("error in the get, ", err)
		return false
	}

	err = json.Unmarshal(body, &session)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Println("error in the unmarshal: ", err, " \nbody: ", string(body))
	}

	return session[0].DateStart.Before(time.Now().UTC()) && session[0].DateEnd.After(time.Now().UTC())
}

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

func TickedDone() string {
	now := time.Now().UTC()
	previus := now.Add(time.Duration(-1) * time.Second)

	Now = strings.ReplaceAll(now.Format("2006-01-02 15:04:05"), " ", "T")
	Previus = strings.ReplaceAll(previus.Format("2006-01-02 15:04:05"), " ", "T")

	return CarFunc()
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

func NoSession() [][]string {
	var positionLastSession []Position
	positionLastSessionUrl := URLSite + "position?session_key=latest"
	body, err := GetData(positionLastSessionUrl)
	if err != nil {
		log.Println("error in the get, ", err)
		return nil
	}

	err = json.Unmarshal(body, &positionLastSession)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Println(string(body))
		return nil
	}

	cleanedSession := cleanSession(positionLastSession)

	return sortSession(cleanedSession)
}

func sortSession(clSe map[int]Position) [][]string {
	var soSession = make([][]string, 20)

	for _, elem := range clSe {
		soSession[elem.Position-1] = []string{
			fmt.Sprintf("%d", elem.Position),
			drvMap[elem.DriverNumber].FirstName,
			drvMap[elem.DriverNumber].LastName,
			fmt.Sprintf("%d", elem.DriverNumber),
			drvMap[elem.DriverNumber].TeamName,
		}
	}

	return soSession
}

func cleanSession(pos []Position) map[int]Position {

	var mapPos = make(map[int]Position)

	for _, elem := range pos {
		value, in := mapPos[elem.DriverNumber]
		if !in {
			mapPos[elem.DriverNumber] = elem
			continue
		}

		if !elem.Date.After(value.Date) {
			mapPos[elem.DriverNumber] = elem
			continue
		}
	}

	return mapPos
}
