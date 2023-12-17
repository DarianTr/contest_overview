package main

import (
	"contest_overview/contest"
	"log"
	"net/http"
)

type Data struct {
	Contest []contest.Contest
	Judges  []string
}

func main() {
	contest.SetJudges()
	contest.SetDomjAPIToken()
	contest.UpdateContests()

	http.HandleFunc("/", h1)
	http.HandleFunc("/search", h2)
	http.HandleFunc("/options", h3)
  http.HandleFunc("/view", view)
	http.HandleFunc("/calendar", displayCalendar)
	http.HandleFunc("/c", displayCalendar)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
