package takedata

import (
	"encoding/json"

	"github.com/mattemello/f1Terminal/internal/errorsh"
)

func takeLaps(previusLap *map[int]Laps) {
	lapUrl := URLSite + "laps?session_key=latest&date_start>" + Previus + "&date_start<=" + Now

	body, err := getData(lapUrl)
	if errorsh.AssertNilFile(err, "The program failed to take the laps data") {
		return
	}

	var lapsAll []LapsAll

	err = json.Unmarshal(body, &lapsAll)
	errorsh.AssertNilJson(err, body)

	controlLaps(lapsAll, previusLap)
}

func controlLaps(lpAll []LapsAll, prLap *map[int]Laps) {
	for _, elem := range lpAll {
		value, ok := (*prLap)[elem.DriverNumber]

		if !ok {
			(*prLap)[elem.DriverNumber] = Laps{
				DurationSector1: elem.DurationSector1,
				DurationSector2: elem.DurationSector2,
				DurationSector3: elem.DurationSector3,
				DateStart:       elem.DateStart,
				DriverNumber:    elem.DriverNumber,
				LapDuration:     elem.LapDuration,
				LapNumber:       elem.LapNumber,
			}
		} else if value.DateStart.Before(elem.DateStart) {
			(*prLap)[elem.DriverNumber] = Laps{
				DurationSector1: elem.DurationSector1,
				DurationSector2: elem.DurationSector2,
				DurationSector3: elem.DurationSector3,
				DateStart:       elem.DateStart,
				DriverNumber:    elem.DriverNumber,
				LapDuration:     elem.LapDuration,
				LapNumber:       elem.LapNumber,
			}
		}
	}
}
