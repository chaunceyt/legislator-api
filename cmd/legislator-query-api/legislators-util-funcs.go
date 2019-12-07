package main

import (
	"io/ioutil"
	"net/http"
)

// loadJSONData - load json file from filesystem.
func loadJSONData(jsonFile string) string {

	jsLegistatorData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	return string(jsLegistatorData)

}

// termType - Convert rep and sen to full terms.
func termType(legType string) string {
	legislatorType := ""
	switch legType {
	case "rep":
		legislatorType = "Representative"
	case "sen":
		legislatorType = "Senator"
	}
	return legislatorType
}

// secureHeaders - send secure headers
func secureHeaders(w http.ResponseWriter) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "same-origin")
	w.Header().Set("Vary", "Accept-Encoding")
	w.WriteHeader(http.StatusOK)
}
