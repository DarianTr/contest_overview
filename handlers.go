package main

import (
	"contest_overview/calendar"
	"contest_overview/contest"
	"fmt"
	"net/http"
	"sort"
	"text/template"
)

var h1 = func(w http.ResponseWriter, r *http.Request) {
	if contest.UpdateNeeded() {
		contest.UpdateContests()
	}
	contests := contest.CONTESTS
	tmpl, _ := template.New("index.html").Funcs(funcMap).ParseFiles("index.html")
	tmpl.Execute(w, Data{
		Contest: contests,
		Judges:  contest.JUDGES,
	})
}

var h2 = func(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("results.html"))
	fmt.Println("htmx resp: ", r.URL.Query().Get("key"))
	tmpl.Execute(w, nil)
}

var h3 = func(w http.ResponseWriter, r *http.Request) {
	if contest.UpdateNeeded() {
		contest.UpdateContests()
	}
	contests := contest.CONTESTS
	var judges []string
	for _, j := range contest.JUDGES {
		if r.URL.Query().Get(j) == "on" {
			judges = append(judges, j)
		}
	}
	contests = contest.FilterContest(contests, contest.FilterForJudge, judges)
	switch r.URL.Query().Get("sorted_by") {
	case "by_date":
		sort.Sort(contest.ByDate(contests))
	case "by_judge":
		sort.Sort(contest.ByJudge(contests))
	default:

	}
	tmpl, _ := template.New("table.html").Funcs(funcMap).ParseFiles("table.html")
	tmpl.Execute(w, Data{
		Contest: contests,
		Judges:  contest.JUDGES,
	})
}

var example = []calendar.Date{
	{Number: 100000000000, IsToday: false},
	{Number: 200000000000, IsToday: false},
	{Number: 300000000000, IsToday: false},
	{Number: 400000000000, IsToday: false},
	{Number: 500000000000, IsToday: false},
	{Number: 600000000000, IsToday: false},
	{Number: 700000000000, IsToday: false},
	{Number: 800000000000, IsToday: true},
	{Number: 900000000000, IsToday: false},
	{Number: 1000000000000, IsToday: false},
}

var weeks = []calendar.Week{
	{Days: example[0:7]},
	{Days: example[7:10]},
}

var displayCalendar = func(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("calendar.html").ParseFiles("./calendar/calendar.html")
	if err != nil {
		fmt.Println(err)
	} else {
		tmpl.Execute(w, weeks)
	}
}

var funcMap = template.FuncMap{
	"Div": func(a int, b int) int {
		return a / b
	},
	"getName": func(c contest.Contest) string {
		return c.GetName()
	},
	"getDate": func(c contest.Contest) string {
		return c.GetDate()
	},
	"getUrl": func(c contest.Contest) string {
		return c.GetUrl()
	},
}
