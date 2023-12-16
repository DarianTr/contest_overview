package main

import (
	"contest_overview/calendar"
	"fmt"
	"net/http"
	"sort"
	"text/template"
)

var h1 = func(w http.ResponseWriter, r *http.Request) {
	if UpdateNeeded() {
		UpdateContests()
	}
	contests := CONTESTS
	tmpl, _ := template.New("index.html").Funcs(funcMap).ParseFiles("index.html")
	tmpl.Execute(w, Data{
		Contest: contests,
		Judges:  JUDGES,
	})
}

var h2 = func(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("results.html"))
	fmt.Println("htmx resp: ", r.URL.Query().Get("key"))
	tmpl.Execute(w, nil)
}

var h3 = func(w http.ResponseWriter, r *http.Request) {
	if UpdateNeeded() {
		UpdateContests()
	}
	contests := CONTESTS
	var judges []string
	for _, j := range JUDGES {
		if r.URL.Query().Get(j) == "on" {
			judges = append(judges, j)
		}
	}

	contests = filter(contests, FilterForJudge, judges)
	switch r.URL.Query().Get("sorted_by") {
	case "by_date":
		sort.Sort(ByDate(contests))
	case "by_judge":
		sort.Sort(ByJudge(contests))
	default:

	}
	tmpl, _ := template.New("table.html").Funcs(funcMap).ParseFiles("table.html")
	tmpl.Execute(w, Data{
		Contest: contests,
		Judges:  JUDGES,
	})
}

var example = []calendar.Date{
	{Number: 1, IsToday: false},
	{Number: 2, IsToday: false},
	{Number: 3, IsToday: false},
	{Number: 4, IsToday: false},
	{Number: 5, IsToday: false},
	{Number: 6, IsToday: false},
	{Number: 7, IsToday: false},
	{Number: 8, IsToday: true},
	{Number: 9, IsToday: false},
	{Number: 10, IsToday: false},
}

var weeks = []calendar.Week{
	{Days: example[0:5]},
	{Days: example[5:10]},
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
	"getName": func(c Contest) string {
		return c.GetName()
	},
	"getDate": func(c Contest) string {
		return c.GetDate()
	},
	"getUrl": func(c Contest) string {
		return c.GetUrl()
	},
}
