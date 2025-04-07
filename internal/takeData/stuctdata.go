package takedata

import "time"

type Session struct {
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

type Driver struct {
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
