package main

// Legislator struct to manage the Legistator json file
// https://mholt.github.io/json-to-go/
// https://theunitedstates.io/congress-legislators/legislators-current.json
type Legislator struct {
	ID struct {
		Bioguide       string   `json:"bioguide"`
		Thomas         string   `json:"thomas"`
		Lis            string   `json:"lis"`
		Govtrack       int      `json:"govtrack"`
		Opensecrets    string   `json:"opensecrets"`
		Votesmart      int      `json:"votesmart"`
		Fec            []string `json:"fec"`
		Cspan          int      `json:"cspan"`
		Wikipedia      string   `json:"wikipedia"`
		HouseHistory   int      `json:"house_history"`
		Ballotpedia    string   `json:"ballotpedia"`
		Maplight       int      `json:"maplight"`
		Icpsr          int      `json:"icpsr"`
		Wikidata       string   `json:"wikidata"`
		GoogleEntityID string   `json:"google_entity_id"`
	} `json:"id"`
	Name struct {
		First        string `json:"first"`
		Last         string `json:"last"`
		NickName     string `json:"nickname"`
		OfficialFull string `json:"official_full"`
	} `json:"name"`
	Bio struct {
		Birthday string `json:"birthday"`
		Gender   string `json:"gender"`
	} `json:"bio"`
	Terms []struct {
		Type        string `json:"type"`
		Start       string `json:"start"`
		End         string `json:"end"`
		State       string `json:"state"`
		District    int    `json:"district,omitempty"`
		Party       string `json:"party"`
		URL         string `json:"url,omitempty"`
		Class       int    `json:"class,omitempty"`
		Address     string `json:"address,omitempty"`
		Phone       string `json:"phone,omitempty"`
		Fax         string `json:"fax,omitempty"`
		ContactForm string `json:"contact_form,omitempty"`
		Office      string `json:"office,omitempty"`
		StateRank   string `json:"state_rank,omitempty"`
		RssURL      string `json:"rss_url,omitempty"`
	} `json:"terms"`
	LeadershipRoles []struct {
		Title   string `json:"title"`
		Chamber string `json:"chamber"`
		Start   string `json:"start"`
	} `json:"leadership_roles,omitempty"`
	OtherNames []struct {
		Last string `json:"last"`
	} `json:"other_names,omitempty"`
	Family []struct {
		Name     string `json:"name"`
		Relation string `json:"relation"`
	} `json:"family,omitempty"`
}
