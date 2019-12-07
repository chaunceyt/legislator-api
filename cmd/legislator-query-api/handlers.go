package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func LegislatorsIndex(w http.ResponseWriter, r *http.Request) {
	secureHeaders(w)

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "Welcome!")
	fmt.Fprintln(w, "My hostname: "+hostname)
	fmt.Fprintln(w, "Your hostname: "+r.RemoteAddr)
}

func Index(w http.ResponseWriter, r *http.Request) {
	secureHeaders(w)

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	data := LegislatorsPageData{
		PageTitle:   "Legislators API",
		PodHostname: hostname,
	}
	tmpl.Execute(w, data)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	secureHeaders(w)

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("about.html"))
	data := LegislatorsPageData{
		PageTitle:   "Legislators API About",
		PodHostname: hostname,
	}
	tmpl.Execute(w, data)
}

// healthz endpoint https://callistaenterprise.se/blogg/teknik/2017/03/22/go-blog-series-part6/
