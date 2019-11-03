package main

import (
	"io/ioutil"
)

func loadJSONData(jsonFile string) string {

	jsLegistatorData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	return string(jsLegistatorData)

}

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
