package takedata

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mattemello/f1Terminal/internal/errorsh"
)

// NOTE: the int in the map is the number of the Driver

func IsSessionOn() (bool, string) {
	var session []SessionStr
	sessionURL := URLSite + "sessions?session_key=latest&meeting_key=latest"

	body, err := getData(sessionURL)
	errorsh.AssertNilTer(err, "The program failed to take the session")

	err = json.Unmarshal(body, &session)
	errorsh.AssertNilJson(err, body)

	return session[0].DateStart.Before(time.Now().UTC()) && session[0].DateEnd.After(time.Now().UTC()), session[0].SessionName
}

func session() map[int]Position {
	var positionLastSession []Position
	positionLastSessionUrl := URLSite + "position?session_key=latest&date<=" + Now
	body, err := getData(positionLastSessionUrl)
	errorsh.AssertNilTer(err, "The program failed to take the session")

	err = json.Unmarshal(body, &positionLastSession)
	errorsh.AssertNilJson(err, body)

	return cleanSession(positionLastSession)
}

func NoSession() [][]string {
	var positionLastSession []Position
	positionLastSessionUrl := URLSite + "position?session_key=latest"
	body, err := getData(positionLastSessionUrl)
	errorsh.AssertNilTer(err, "The program failed to take the session")

	err = json.Unmarshal(body, &positionLastSession)
	errorsh.AssertNilJson(err, body)

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

		if elem.Date.After(value.Date) {
			mapPos[elem.DriverNumber] = elem
			continue
		}
	}

	return mapPos
}
