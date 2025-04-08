package takedata

import (
	"encoding/json"
	"log"
)

// NOTE: the int in the map is the number of the Driver
func interval() map[int]IntervalAll {
	var inter []IntervalAll
	intUrl := URLSite + "intervals?session_key=latest&date>" + Previus + "&date<=" + Now

	body, err := getData(intUrl)
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
