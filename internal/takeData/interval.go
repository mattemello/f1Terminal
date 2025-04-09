package takedata

import (
	"encoding/json"

	"github.com/mattemello/f1Terminal/internal/errorsh"
)

// NOTE: the int in the map is the number of the Driver
func interval() map[int]IntervalAll {
	var inter []IntervalAll
	intUrl := URLSite + "intervals?session_key=latest&date>" + Previus + "&date<=" + Now

	body, err := getData(intUrl)
	errorsh.AssertNilFile(err, "The program failed to get the data")

	err = json.Unmarshal(body, &inter)
	errorsh.AssertNilJson(err, body)

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
