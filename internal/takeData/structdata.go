package takedata

import "time"

type LapsAll struct {
	DateStart       time.Time `json:"date_start"`
	DriverNumber    int       `json:"driver_number"`
	DurationSector1 float64   `json:"duration_sector_1"`
	DurationSector2 float64   `json:"duration_sector_2"`
	DurationSector3 float64   `json:"duration_sector_3"`
	I1Speed         int       `json:"i1_speed"`
	I2Speed         int       `json:"i2_speed"`
	IsPitOutLap     bool      `json:"is_pit_out_lap"`
	LapDuration     float64   `json:"lap_duration"`
	LapNumber       int       `json:"lap_number"`
	MeetingKey      int       `json:"meeting_key"`
	SegmentsSector1 []int     `json:"segments_sector_1"`
	SegmentsSector2 []int     `json:"segments_sector_2"`
	SegmentsSector3 []int     `json:"segments_sector_3"`
	SessionKey      int       `json:"session_key"`
	StSpeed         int       `json:"st_speed"`
}

type Laps struct {
	DateStart       time.Time
	DriverNumber    int
	DurationSector1 float64
	DurationSector2 float64
	DurationSector3 float64
	LapDuration     float64
	LapNumber       int
}

type SessionStr struct {
	CircuitKey       int       `json:"circuit_key"`
	CircuitShortName string    `json:"circuit_short_name"`
	CountryCode      string    `json:"country_code"`
	CountryKey       int       `json:"country_key"`
	CountryName      string    `json:"country_name"`
	DateEnd          time.Time `json:"date_end"`
	DateStart        time.Time `json:"date_start"`
	GmtOffset        string    `json:"gmt_offset"`
	Location         string    `json:"location"`
	MeetingKey       int       `json:"meeting_key"`
	SessionKey       int       `json:"session_key"`
	SessionName      string    `json:"session_name"`
	SessionType      string    `json:"session_type"`
	Year             int       `json:"year"`
}

type Circuit struct {
	CircuitKey          int       `json:"circuit_key"`
	CircuitShortName    string    `json:"circuit_short_name"`
	CountryCode         string    `json:"country_code"`
	CountryKey          int       `json:"country_key"`
	CountryName         string    `json:"country_name"`
	DateStart           time.Time `json:"date_start"`
	GmtOffset           string    `json:"gmt_offset"`
	Location            string    `json:"location"`
	MeetingKey          int       `json:"meeting_key"`
	MeetingName         string    `json:"meeting_name"`
	MeetingOfficialName string    `json:"meeting_official_name"`
	Year                int       `json:"year"`
}

type DriverAll struct {
	BroadcastName string `json:"broadcast_name"`
	CountryCode   string `json:"country_code"`
	DriverNumber  int    `json:"driver_number"`
	FirstName     string `json:"first_name"`
	FullName      string `json:"full_name"`
	HeadshotURL   string `json:"headshot_url"`
	LastName      string `json:"last_name"`
	MeetingKey    int    `json:"meeting_key"`
	NameAcronym   string `json:"name_acronym"`
	SessionKey    int    `json:"session_key"`
	TeamColour    string `json:"team_colour"`
	TeamName      string `json:"team_name"`
}

type Driver struct {
	DriverNumber int
	FirstName    string
	LastName     string
	NameAcronym  string
	TeamName     string
}

type IntervalAll struct {
	Date         time.Time `json:"date"`
	DriverNumber int       `json:"driver_number"`
	GapToLeader  float64   `json:"gap_to_leader"`
	Interval     float64   `json:"interval"`
	MeetingKey   int       `json:"meeting_key"`
	SessionKey   int       `json:"session_key"`
}

type Interval struct {
	DriverNumber int
	GapToLeader  string
	Interval     string
	Date         time.Time
}

type Car struct {
	URL     string
	CarData CarData
}

type CarData []struct {
	Brake        int       `json:"brake"`
	Date         time.Time `json:"date"`
	DriverNumber int       `json:"driver_number"`
	Drs          int       `json:"drs"`
	MeetingKey   int       `json:"meeting_key"`
	NGear        int       `json:"n_gear"`
	Rpm          int       `json:"rpm"`
	SessionKey   int       `json:"session_key"`
	Speed        int       `json:"speed"`
	Throttle     int       `json:"throttle"`
}

type Position struct {
	Date         time.Time `json:"date"`
	DriverNumber int       `json:"driver_number"`
	MeetingKey   int       `json:"meeting_key"`
	Position     int       `json:"position"`
	SessionKey   int       `json:"session_key"`
}
