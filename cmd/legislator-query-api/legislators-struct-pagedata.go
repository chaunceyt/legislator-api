package main

type LegislatorDistrictOffice struct {
	ID             string
	Address        string
	Suite          string
	City           string
	State          string
	Zip            string
	Latitude       float64
	Longitude      float64
	Phone          string
	Fax            string
	Building       string
	LegislatorName string
}

type LegislatorsPageData struct {
	PageTitle                    string
	PodHostname                  string
	LegislatorFirstName          string
	LegislatorLastName           string
	LegislatorName               string
	LegislatorMailingAddress     string
	LegislatorOfficePhone        string
	LegislatorParty              string
	LegislatorOfficeFax          string
	LegislatorWebsite            string
	LegislatorBioguideId         string
	LegislatorBioguideUrl        string
	LegislatorWikiDataUrl        string
	LegislatorVoteSmartUrl       string
	LegislatorGovTrackUrl        string
	LegislatorFacebookUrl        string
	LegislatorTwitterUrl         string
	LegislatorYoutubeChannelUrl  string
	LegislatorOpenSecrets        string
	LegislatorHouseHistoryUrl    string
	LegislatorMaplightProfileUrl string
	LegislatorContactFormUrl     string
	LegislatorRssFeed            string
	LegislatorDistrictOffices    []LegislatorDistrictOffice
}
