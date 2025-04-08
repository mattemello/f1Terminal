package takedata

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const URLSite = "https://api.openf1.org/v1/"

var Now string
var Previus string

func getData(url string) ([]byte, error) {
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

func TakeCircuit() (Circuit, error) {
	var cir []Circuit
	circuitUrl := URLSite + "meetings?meeting_key=latest"

	body, err := getData(circuitUrl)
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

func TickedDone() [][]string {
	now := time.Now().UTC()
	previus := now.Add(time.Duration(-1) * time.Second)

	Now = strings.ReplaceAll(now.Format("2006-01-02 15:04:05"), " ", "T")
	Previus = strings.ReplaceAll(previus.Format("2006-01-02 15:04:05"), " ", "T")

	session := session()
	inter := interval()

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
