package takedata

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mattemello/f1Terminal/internal/errorsh"
)

var drvMap = make(map[int]Driver)

func TakeDriverInSession() {
	var driver []DriverAll
	drivUrl := URLSite + "drivers?session_key=latest"

	body, err := getData(drivUrl)
	errorsh.AssertNilTer(err, "The program failed to take the drivers data")

	err = json.Unmarshal(body, &driver)
	if errorsh.AssertNilJson(err, body) {
		fmt.Println("Error in the parse of the driver, to see more please control the log file")
		os.Exit(1)
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
