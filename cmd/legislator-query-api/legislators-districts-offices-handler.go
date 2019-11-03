package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// LegislatorsInfo return district offices
func LegislatorsInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	firstname := vars["firstname"]
	lastname := vars["lastname"]
	jsLegislatorsData := loadJSONData(currentLegislatorsJSONFile)
	bsLegislatorsData := []byte(jsLegislatorsData)

	jsLegislatorsSocialMediaData := loadJSONData(currentLegislatorsSocialMediaJSONFile)
	bsLegislatorsSocialMediaData := []byte(jsLegislatorsSocialMediaData)

	jsLegislatorsDistrictOfficesData := loadJSONData(currentLegislatorsDistrictOfficesJSONFile)
	bsLegislatorsDistrictOfficesData := []byte(jsLegislatorsDistrictOfficesData)

	//legislator := []Legislator{}
	var legislator []Legislator
	var socialMedia []LegislatorsSocialMedia
	var officesData []LegislatorsDistrictOffices

	err := json.Unmarshal(bsLegislatorsData, &legislator)
	if err != nil {
		println(err)
	}

	err = json.Unmarshal(bsLegislatorsSocialMediaData, &socialMedia)
	if err != nil {
		println(err)
	}

	err = json.Unmarshal(bsLegislatorsDistrictOfficesData, &officesData)
	if err != nil {
		println(err)
	}

	for _, v := range legislator {

		// Search for firstname or nickname
		if (v.Name.First == strings.Title(firstname) || v.Name.NickName == strings.Title(firstname)) && v.Name.Last == strings.Title(lastname) {

			// Get current term in office. This is the last item in the array.
			// A type:rep has a district the sen doesn't

			fmt.Fprintln(w, "District Offices\n")

			for _, o := range officesData {
				if o.ID.Bioguide == v.ID.Bioguide {
					for _, od := range o.Offices {
						officeID := strings.Replace(od.ID, o.ID.Bioguide+"-", "", -1)
						fmt.Fprintln(w, strings.Title(officeID)+" District Office \n\n", od.Address+" "+od.Suite+" "+od.Building+" \n "+od.City+" "+od.State+" "+od.Zip+" \n Phone: "+od.Phone+" \n Fax: "+od.Fax+" \n")
					}
				}
			}
		}
	}
}
