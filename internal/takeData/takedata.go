package takedata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	if errorsh.AssertNilJson(err, body) {
		fmt.Println("Error in the parse of the driver, to see more please control the log file")
		os.Exit(1)
	}

	return cir[0]
}

func TickedDone(lap *map[int]Laps) [][]string {
	now := time.Now().UTC()
	previus := now.Add(time.Duration(-1) * time.Second)

	Now = strings.ReplaceAll(now.Format("2006-01-02 15:04:05"), " ", "T")
	Previus = strings.ReplaceAll(previus.Format("2006-01-02 15:04:05"), " ", "T")

	session := session()
	if len(session) == 0 {
		return nil
	}
	//inter := interval()
	takeLaps(lap)

	return changedTable(session, lap)
}

func changedTable(clSe map[int]Position, lap *map[int]Laps) [][]string {
	var driv = make([][]string, 20)

	for _, elem := range clSe {
		driv[elem.Position-1] = []string{
			fmt.Sprintf("%d", elem.Position),
			drvMap[elem.DriverNumber].FirstName,
			drvMap[elem.DriverNumber].LastName,
			fmt.Sprintf("%d", elem.DriverNumber),
			fmt.Sprintf("%f", (*lap)[elem.DriverNumber].LapDuration),
			fmt.Sprintf("%d", (*lap)[elem.DriverNumber].LapNumber),
			drvMap[elem.DriverNumber].TeamName,
		}
	}

	// for _, elem := range driv {
	// 	fmt.Println(elem)
	// }

	return driv
}
