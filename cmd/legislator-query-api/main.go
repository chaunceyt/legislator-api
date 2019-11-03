package main

import (
	"log"
	"net/http"
)

const currentLegislatorsSocialMediaJSONFile = "data/legislators-social-media.json"
const currentLegislatorsDistrictOfficesJSONFile = "data/legislators-district-offices.json"
const currentLegislatorsJSONFile = "data/legislators-current.json"
const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func main() {

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
