package takedata

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// NOTE: the int in the map is the number of the Driver

func IsSessionOn() bool {
	var session []SessionStr
	sessionURL := URLSite + "sessions?session_key=latest&meeting_key=latest"

	body, err := getData(sessionURL)

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

func session() map[int]Position {
	var positionLastSession []Position
	positionLastSessionUrl := URLSite + "position?session_key=latest&date>" + Previus + "&date<=" + Now
	body, err := getData(positionLastSessionUrl)
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

	return cleanSession(positionLastSession)
}

func NoSession() [][]string {
	var positionLastSession []Position
	positionLastSessionUrl := URLSite + "position?session_key=latest"
	body, err := getData(positionLastSessionUrl)
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
