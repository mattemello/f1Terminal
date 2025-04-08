package takedata

import (
	"encoding/json"
	"log"
)

var drvMap = make(map[int]Driver)

func TakeDriverInSession() {
	var driver []DriverAll
	drivUrl := URLSite + "drivers?session_key=latest"

	body, err := getData(drivUrl)
	if err != nil {
		log.Println("error in the get, ", err)
		return
	}

	err = json.Unmarshal(body, &driver)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Println("error in the unmarshal: ", err, " \nbody: ", string(body))
	}

	driverMap(driver)
}

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
