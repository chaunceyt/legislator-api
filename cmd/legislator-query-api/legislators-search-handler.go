package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// LegislatorsSearch - search for a legislator.
func LegislatorsSearch(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")

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

	var lawmaker = make(map[string]string)
	for _, v := range legislator {

		// Search for firstname or nickname
		if (v.Name.First == strings.Title(firstname) || v.Name.NickName == strings.Title(firstname)) && v.Name.Last == strings.Title(lastname) {

			legislatorName := ""
			if v.Terms[len(v.Terms)-1].Type == "rep" {
				legislatorName = termType(v.Terms[len(v.Terms)-1].Type) + " " + v.Name.OfficialFull + " (" + v.Terms[len(v.Terms)-1].Party + ") " + v.Terms[len(v.Terms)-1].State + " District " + strconv.Itoa(v.Terms[len(v.Terms)-1].District)
			} else {
				legislatorName = termType(v.Terms[len(v.Terms)-1].Type) + " " + v.Name.OfficialFull + " (" + v.Terms[len(v.Terms)-1].Party + ") " + v.Terms[len(v.Terms)-1].State
			}

			lawmaker["bioguide_id"] = v.ID.Bioguide
			lawmaker["name"] = legislatorName
			lawmaker["mailing_address"] = v.Terms[len(v.Terms)-1].Address
			lawmaker["office_phone"] = v.Terms[len(v.Terms)-1].Phone
			lawmaker["party"] = v.Terms[len(v.Terms)-1].Party

			// Fax number.
			if v.Terms[len(v.Terms)-1].Fax != "" {
				lawmaker["office_fax"] = v.Terms[len(v.Terms)-1].Fax
			} else {
				lawmaker["office_fax"] = ""
			}

			lawmaker["website"] = v.Terms[len(v.Terms)-1].URL

			// Handle social media
			for _, s := range socialMedia {
				if s.ID.Bioguide == v.ID.Bioguide {
					if s.Social.Facebook != "" {
						lawmaker["facebook"] = "https://www.facebook.com/" + s.Social.Facebook
					} else {
						lawmaker["facebook"] = ""
					}

					if s.Social.Twitter != "" {
						lawmaker["twitter"] = "https://twitter.com/" + s.Social.Twitter
					} else {
						lawmaker["twitter"] = ""
					}

					if s.Social.YoutubeID != "" {
						lawmaker["youtube_channel"] = "https://www.youtube.com/channel/" + s.Social.YoutubeID
					} else {
						lawmaker["youtube_channel"] = ""
					}

				}

			}

			lawmaker["biography"] = "http://bioguide.congress.gov/scripts/biodisplay.pl?index=" + v.ID.Bioguide
			//lawmaker["wikidata_summary"] = "http://bioguide.congress.gov/scripts/biodisplay.pl?index=" + v.ID.Bioguide
			lawmaker["votesmart_profile"] = "https://votesmart.org/candidate/" + strconv.Itoa(v.ID.Votesmart)
			lawmaker["govtrack_profile"] = "https://www.govtrack.us/congress/members/" + strconv.Itoa(v.ID.Govtrack)

			if v.ID.Opensecrets != "" {
				lawmaker["opensecrets"] = "https://www.opensecrets.org/members-of-congress/summary?cid=" + v.ID.Opensecrets
			} else {
				lawmaker["opensecrets"] = ""
			}

			if v.ID.HouseHistory != 0 {
				lawmaker["house_history"] = "https://history.house.gov/People/Detail/" + strconv.Itoa(v.ID.HouseHistory)
			} else {
				lawmaker["house_history"] = ""
			}

			if v.ID.Maplight != 0 {
				lawmaker["maplight_profile"] = "http://classic.maplight.org/us-congress/legislator/" + strconv.Itoa(v.ID.Maplight)
			} else {
				lawmaker["maplight_profile"] = ""
			}

			if v.Terms[len(v.Terms)-1].RssURL != "" {
				lawmaker["contact_form"] = v.Terms[len(v.Terms)-1].ContactForm
			} else {
				lawmaker["contact_form"] = ""
			}

			if v.Terms[len(v.Terms)-1].RssURL != "" {
				lawmaker["rssfeed"] = v.Terms[len(v.Terms)-1].RssURL
			} else {
				lawmaker["rssfeed"] = ""
			}

			var legislatorOfficeList []LegislatorDistrictOffice
			for _, o := range officesData {
				if o.ID.Bioguide == v.ID.Bioguide {
					for _, od := range o.Offices {
						officeID := strings.Replace(od.ID, o.ID.Bioguide+"-", "", -1)
						legislatorOfficeList = append(legislatorOfficeList, LegislatorDistrictOffice{
							LegislatorName: firstname + " " + lastname,
							ID:             strings.Title(strings.Replace(officeID, "_", " ", -1)),
							Address:        od.Address,
							Suite:          od.Suite,
							Building:       od.Building,
							City:           od.City,
							State:          od.State,
							Zip:            od.Zip,
							Phone:          od.Phone,
							Fax:            od.Fax,
							Longitude:      od.Longitude,
							Latitude:       od.Latitude,
						})
					}
				}
			}
			tmpl := template.Must(template.ParseFiles("home.html"))
			data := LegislatorsPageData{
				PageTitle:                    "Legislators API",
				LegislatorName:               lawmaker["name"],
				LegislatorFirstName:          firstname,
				LegislatorLastName:           lastname,
				LegislatorMailingAddress:     lawmaker["mailing_address"],
				LegislatorOfficePhone:        lawmaker["office_phone"],
				LegislatorParty:              lawmaker["party"],
				LegislatorOfficeFax:          lawmaker["office_fax"],
				LegislatorWebsite:            lawmaker["website"],
				LegislatorBioguideUrl:        lawmaker["biography"],
				LegislatorVoteSmartUrl:       lawmaker["votesmart_profile"],
				LegislatorGovTrackUrl:        lawmaker["govtrack_profile"],
				LegislatorBioguideId:         lawmaker["bioguide_id"],
				LegislatorFacebookUrl:        lawmaker["facebook"],
				LegislatorTwitterUrl:         lawmaker["twitter"],
				LegislatorYoutubeChannelUrl:  lawmaker["youtube_channel"],
				LegislatorOpenSecrets:        lawmaker["opensecrets"],
				LegislatorHouseHistoryUrl:    lawmaker["house_history"],
				LegislatorMaplightProfileUrl: lawmaker["maplight_profile"],
				LegislatorContactFormUrl:     lawmaker["contact_form"],
				LegislatorRssFeed:            lawmaker["rssfeed"],
				LegislatorDistrictOffices:    legislatorOfficeList,
			}
			tmpl.Execute(w, data)
		}
	}
}
