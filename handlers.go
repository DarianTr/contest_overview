package main

import (
	"fmt"
	"net/http"
	"sort"
	"text/template"
)

var h1 = func(w http.ResponseWriter, r *http.Request) {
	var contests []Contest
	contests = append(contests, filter(ToContests(GetCodeforces().Result), FilterIsUpcoming, nil)...)
	contests = append(contests, filter(DmojToContests(GetDmoj().Data.Objects), FilterIsUpcoming, nil)...)
	contests = append(contests, GetAtCoder()...)
	// if r.Method == "POST" {
	// 	r.ParseForm()
	// 	s := r.Form["sorted_by"]
	// 	if len(s) > 0 {
	// 		if s[0] == "by_date" {
	// 			sort.Sort(ByDate(contests))
	// 		} else if s[0] == "by_judge" {
	// 			sort.Sort(ByJudge(contests))
	// 		}
	// 		fmt.Println(s[0])
	// 	}
	// 	fmt.Println(reflect.TypeOf(r.Form["Codeforces"]))
	// 	var judges []string
	// 	codeforces := r.Form["Codeforces"]
	// 	dmoj := r.Form["Dmoj"]
	// 	if len(codeforces) > 0 && codeforces[0] == "on" {
	// 		judges = append(judges, "Codeforces")
	// 		fmt.Println("dmoj", judges)
	// 	}
	// 	if len(dmoj) > 0 && dmoj[0] == "on" {
	// 		judges = append(judges, "Dmoj")
	// 		fmt.Println("dmoj", judges)
	// 	}
	// 	contests = filter(contests, FilterForJudge, judges)

	// }
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

	fmt.Println(r.URL)

	var contests []Contest
	var judges []string
	contests = append(contests, filter(ToContests(GetCodeforces().Result), FilterIsUpcoming, nil)...)
	contests = append(contests, filter(DmojToContests(GetDmoj().Data.Objects), FilterIsUpcoming, nil)...)
	contests = append(contests, GetAtCoder()...)

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
