package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// LegislatorsJSON return a json object for the congressional member.
func LegislatorsJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	firstname := vars["firstname"]
	lastname := vars["lastname"]
	jsLegislatorsData := loadJSONData(currentLegislatorsJSONFile)
	bsLegislatorsData := []byte(jsLegislatorsData)

	jsLegislatorsSocialMediaData := loadJSONData(currentLegislatorsSocialMediaJSONFile)
	bsLegislatorsSocialMediaData := []byte(jsLegislatorsSocialMediaData)

	jsLegislatorsDistrictOfficesData := loadJSONData(currentLegislatorsDistrictOfficesJSONFile)
	bsLegislatorsDistrictOfficesData := []byte(jsLegislatorsDistrictOfficesData)

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

	var lawmaker = make(map[string]string)
	for _, v := range legislator {

		// Search for firstname or nickname
		if (v.Name.First == strings.Title(firstname) || v.Name.NickName == strings.Title(firstname)) && v.Name.Last == strings.Title(lastname) {

			// Get current term in office. This is the last item in the array.
			// A type:rep has a district the sen doesn't

			legislatorName := ""
			if v.Terms[len(v.Terms)-1].Type == "rep" {
				legislatorName = termType(v.Terms[len(v.Terms)-1].Type) + " " + v.Name.OfficialFull + " (" + v.Terms[len(v.Terms)-1].Party + ") " + v.Terms[len(v.Terms)-1].State + " District " + strconv.Itoa(v.Terms[len(v.Terms)-1].District)
			} else {
				legislatorName = termType(v.Terms[len(v.Terms)-1].Type) + " " + v.Name.OfficialFull + " (" + v.Terms[len(v.Terms)-1].Party + ") " + v.Terms[len(v.Terms)-1].State
			}
			lawmaker["name"] = legislatorName
			lawmaker["mailing_address"] = v.Terms[len(v.Terms)-1].Address
			lawmaker["office_phone"] = v.Terms[len(v.Terms)-1].Phone
			// Fax number.
			if v.Terms[len(v.Terms)-1].Fax != "" {
				lawmaker["office_fax"] = v.Terms[len(v.Terms)-1].Fax
			}

			lawmaker["website"] = v.Terms[len(v.Terms)-1].URL

			// Handle social media
			for _, s := range socialMedia {
				if s.ID.Bioguide == v.ID.Bioguide {
					if s.Social.Facebook != "" {
						lawmaker["facebook"] = "https://www.facebook.com/" + s.Social.Facebook
					}

					if s.Social.Twitter != "" {
						lawmaker["twitter"] = "https://twitter.com/" + s.Social.Twitter
					}

					if s.Social.YoutubeID != "" {
						lawmaker["youtube_channel"] = "https://www.youtube.com/channel/" + s.Social.YoutubeID
					}

				}

			}

			lawmaker["biography"] = "http://bioguide.congress.gov/scripts/biodisplay.pl?index=" + v.ID.Bioguide
			lawmaker["wikidata_summary"] = "http://bioguide.congress.gov/scripts/biodisplay.pl?index=" + v.ID.Bioguide
			lawmaker["votesmart_profile"] = "https://votesmart.org/candidate/" + strconv.Itoa(v.ID.Votesmart)
			lawmaker["govtrack_profile"] = "https://www.govtrack.us/congress/members/" + strconv.Itoa(v.ID.Govtrack)

			if v.ID.Opensecrets != "" {
				lawmaker["opensecrets"] = "https://www.opensecrets.org/members-of-congress/summary?cid=" + v.ID.Opensecrets
			}

			if v.ID.HouseHistory != 0 {
				lawmaker["house_history"] = "https://history.house.gov/People/Detail/" + strconv.Itoa(v.ID.HouseHistory)
			}

			if v.ID.Maplight != 0 {
				lawmaker["maplight_profile"] = "http://classic.maplight.org/us-congress/legislator/" + strconv.Itoa(v.ID.Maplight)
			}

			if v.Terms[len(v.Terms)-1].RssURL != "" {
				lawmaker["contact_form"] = v.Terms[len(v.Terms)-1].ContactForm
			}

			if v.Terms[len(v.Terms)-1].RssURL != "" {
				lawmaker["rssfeed"] = v.Terms[len(v.Terms)-1].RssURL
			}

			// @TODO: Convert this to an array so we can better mange district offices.
			for _, o := range officesData {
				if o.ID.Bioguide == v.ID.Bioguide {
					for _, od := range o.Offices {
						officeID := strings.Replace(od.ID, o.ID.Bioguide+"-", "", -1)
						lawmaker["district_office_"+officeID] = "" + od.Address + " " + od.Suite + " " + od.Building + " " + od.City + " " + od.State + " " + od.Zip + " Phone: " + od.Phone + " Fax: " + od.Fax + ""
					}
				}
			}

			// indent the json object for better render.
			legislatorJSON, err := json.MarshalIndent(lawmaker, "", "   ")
			if err != nil {
				log.Fatal("Cannot encode to JSON ", err)
			}
			fmt.Fprintf(w, "%s", legislatorJSON)

		}
	}
}
