package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// LegislatorByState struct
type LegislatorByState struct {
	gorm.Model
	LastName       string
	FirstName      string
	MiddleName     string
	Suffix         string
	NickName       string
	FullName       string
	Birthday       string
	Gender         string
	Type           string
	State          string
	District       string
	SenateClass    string
	Party          string
	URL            string
	Address        string
	Phone          string
	ContactForm    string
	RssURL         string
	Twitter        string
	Facebook       string
	Youtube        string
	YoutubeID      string
	BioguideID     string
	ThomasID       string
	OpensecretsID  string
	LisID          string
	FecIds         string
	CspanID        string
	GovtrackID     string
	VotesmartID    string
	BalllotpediaID string
	IcpsrID        string
	WikipediaID    string
}

type LegislatorByStatePageData struct {
	PageTitle   string
	Legislators []LegislatorByState
}

// LegislatorsByState - return legislators by state.
func LegislatorsByState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	state := strings.ToUpper(vars["state"])

	db, err := gorm.Open("sqlite3", "data/legislator.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	var legislators []LegislatorByState
	//db.Find(&legislators)
	db.Raw("SELECT * FROM legislators WHERE state = ?", state).Scan(&legislators)
	//json.NewEncoder(w).Encode(legislators)

	tmpl := template.Must(template.ParseFiles("state.html"))
	data := LegislatorByStatePageData{
		PageTitle:   "Legislators By State",
		Legislators: legislators,
	}
	// log.Println(data)
	tmpl.Execute(w, data)

}
