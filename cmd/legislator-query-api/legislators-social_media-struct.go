package main

// LegislatorsSocialMedia struct
// Json file: https://theunitedstates.io/congress-legislators/legislators-social-media.json
type LegislatorsSocialMedia struct {
	ID struct {
		Bioguide string `json:"bioguide"`
		Thomas   string `json:"thomas"`
		Govtrack int    `json:"govtrack"`
	} `json:"id"`
	Social struct {
		Twitter   string `json:"twitter"`
		Facebook  string `json:"facebook"`
		YoutubeID string `json:"youtube_id"`
		TwitterID string `json:"twitter_id"`
	} `json:"social"`
}
