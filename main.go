package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/joho/godotenv"
)

var client *http.Client
var domjAPIToken string

func GetDmoj() DmojResponse {
	var dmoj DmojResponse
	url := "https://dmoj.ca/api/v2/contests"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", domjAPIToken)
	res, err := client.Do(req)
	if err != nil {
		println(err)
	} else {
		defer res.Body.Close()
		json.NewDecoder(res.Body).Decode(&dmoj)
		fmt.Println("res: ", dmoj.ApiVersion)
	}
	return dmoj
}

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

func SetDomjAPIToken() {
	envFile, _ := godotenv.Read(".env")
	domjAPIToken = envFile["DMOJ_API_TOKEN"]

}

type Data struct {
	Data []Contest
}

func main() {
	var _ Contest = DmojContest{}
	var _ Contest = CodeforcesContest{}
	SetDomjAPIToken()
	client = &http.Client{Timeout: 10 * time.Second}

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
		var contests []Contest
		contests = append(contests, filter(to_contests(GetCodeforces().Result), Filter{func(c Contest) bool { return c.is_active() }})...)
		contests = append(contests, filter(dmoj_to_contests(GetDmoj().Data.Objects), Filter{func(c Contest) bool { return c.is_active() }})...)
		if r.Method == "POST" {
			r.ParseForm()
			ans := r.Form["sorted_by"]
			if ans[0] == "by_date" {
				sort.Sort(ByDate(contests))
			} else if ans[0] == "by_judge" {
				sort.Sort(ByJudge(contests))
			}
			fmt.Println(ans[0])
		}
		tmpl, _ := template.New("index.html").Funcs(funcMap).ParseFiles("index.html")
		tmpl.Execute(w, contests)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
