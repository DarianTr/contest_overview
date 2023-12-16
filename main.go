package main

import (
	"contest_overview/contest"
	"log"
	"net/http"
)

var client *http.Client
var domjAPIToken string

type Data struct {
	Contest []contest.Contest
	Judges  []string
}

func main() {
	// SetJudges()
	// SetDomjAPIToken()
	// UpdateContests()

	// http.HandleFunc("/", h1)
	// http.HandleFunc("/search", h2)
	// http.HandleFunc("/options", h3)
	http.HandleFunc("/calendar", displayCalendar)
	http.HandleFunc("/c", displayCalendar)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
