package main

import (
	"contest_overview/calendar"
	"contest_overview/contest"
	"fmt"
	"net/http"
	"sort"
	"text/template"
	"time"
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

var view = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmx resp: ", r.URL.Query().Get("switch"))
}

var displayCalendar = func(w http.ResponseWriter, r *http.Request) {
	data := calendar.DisplayCalendar(calendar.ContestsToCalendar(contest.CONTESTS), 2023, time.Now().Month())
	tmpl, err := template.New("calendar.html").Funcs(funcMap).ParseFiles("./calendar/calendar.html")
	if err != nil {
		fmt.Println(err)
	} else {
		tmpl.Execute(w, data)
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
	"isSunday": func(a int) bool {
		return a%7 == 0
	},
}
