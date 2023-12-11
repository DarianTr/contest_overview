package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

var client *http.Client

func GetCodeforces() CodeforcesResponse {
	url := "https://codeforces.com/api/contest.list?gym=false"
	var codeforces CodeforcesResponse

	err := GetJson(url, &codeforces)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	return codeforces
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

type Data struct {
	Data []Contest
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	var _ Contest = new(CodeforcesContest)
	funcMap := template.FuncMap{
		"Div": func(a int, b int) int {
			return a / b
		},
		"getName": func(c Contest) string {
			return c.get_name()
		},
		"getDate": func(c Contest) string {
			return c.get_date()
		},
		"getUrl": func(c Contest) string {
			return c.get_url()
		},
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		codeforcesResponse := GetCodeforces()
		c := filter(to_contests(codeforcesResponse.Result), Filter{func(c Contest) bool { return c.is_active() }})
		println(len(c))
		sort.Sort(ByDate(c))
		for _, s := range c {
			println(s.get_name())
		}
		tmpl, _ := template.New("index.html").Funcs(funcMap).ParseFiles("index.html")
		tmpl.Execute(w, c)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
