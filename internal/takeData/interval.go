package takedata

import (
	"encoding/json"
	"fmt"

	"github.com/mattemello/f1Terminal/internal/errorsh"
)

// NOTE: the int in the map is the number of the Driver
func interval() map[int]Interval {
	var inter []IntervalAll
	intUrl := URLSite + "intervals?session_key=latest&date>" + Previus + "&date<=" + Now

	body, err := getData(intUrl)
	if errorsh.AssertNilFile(err, "The program failed to get the data") {
		return errorDataInterval()
	}

	err = json.Unmarshal(body, &inter)
	errorsh.AssertNilJson(err, body)

	return cleanInterval(inter)
}

func errorDataInterval() map[int]Interval {
	var intervalMap = make(map[int]Interval)

	for i, _ := range drvMap {
		intervalMap[i] = Interval{
			DriverNumber: i,
			Interval:     "-",
			GapToLeader:  "-",
		}
	}

	return intervalMap
}

func cleanInterval(interval []IntervalAll) map[int]Interval {
	var intervalMap = make(map[int]Interval)

	for _, elem := range interval {
		value, in := intervalMap[elem.DriverNumber]
		if !in {
			intervalMap[elem.DriverNumber] = Interval{
				DriverNumber: elem.DriverNumber,
				Interval:     fmt.Sprintf("%f", elem.Interval),
				GapToLeader:  fmt.Sprintf("%f", elem.GapToLeader),
				Date:         elem.Date,
			}
			continue
		}

		if !elem.Date.After(value.Date) {
			intervalMap[elem.DriverNumber] = Interval{
				DriverNumber: elem.DriverNumber,
				Interval:     fmt.Sprintf("%f", elem.Interval),
				GapToLeader:  fmt.Sprintf("%f", elem.GapToLeader),
				Date:         elem.Date,
			}
			continue
		}
	}

	return intervalMap
}
