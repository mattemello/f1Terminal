package takedata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/mattemello/f1Terminal/internal/errorsh"
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

func TakeCircuit() Circuit {
	var cir []Circuit
	circuitUrl := URLSite + "meetings?meeting_key=latest"

	body, err := getData(circuitUrl)
	errorsh.AssertNilTer(err, "The program failed to get the data")

	err = json.Unmarshal(body, &cir)
	errorsh.AssertNilJson(err, body)

	return cir[0]
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

func changedTable(clSe map[int]Position, inte map[int]Interval) [][]string {
	var driv = make([][]string, 20)

	for _, elem := range clSe {
		driv[elem.Position-1] = []string{
			fmt.Sprintf("%d", elem.Position),
			drvMap[elem.DriverNumber].FirstName,
			drvMap[elem.DriverNumber].LastName,
			fmt.Sprintf("%d", elem.DriverNumber),
			inte[elem.DriverNumber].GapToLeader,
			inte[elem.DriverNumber].Interval,
			drvMap[elem.DriverNumber].TeamName,
		}
	}

	return nil
}
