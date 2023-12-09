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

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	funcMap := template.FuncMap{
		"Div": func(a int, b int) int {
			return a / b
		},
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		codeforcesResponse := GetCodeforces()
		c := filter(codeforcesResponse.Result, Filter{func(c CodeforcesContest) bool { return c.Phase == "BEFORE" }})
		sort.Sort(ByDate(c))
		tmpl, _ := template.New("index.html").Funcs(funcMap).ParseFiles("index.html")
		contests := map[string][]CodeforcesContest{
			"Test": c,
		}
		tmpl.Execute(w, contests)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
