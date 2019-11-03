package main

// LegislatorsDistrictOffices struct.
type LegislatorsDistrictOffices struct {
	ID struct {
		Bioguide string `json:"bioguide"`
		Govtrack int    `json:"govtrack"`
		Thomas   string `json:"thomas"`
	} `json:"id"`
	Offices []struct {
		ID        string  `json:"id"`
		Address   string  `json:"address"`
		Suite     string  `json:"suite"`
		City      string  `json:"city"`
		State     string  `json:"state"`
		Zip       string  `json:"zip"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Phone     string  `json:"phone"`
		Fax       string  `json:"fax"`
		Building  string  `json:"building,omitempty"`
	} `json:"offices"`
}
