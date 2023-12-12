package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	http.HandleFunc("/", h1)
	http.HandleFunc("/search", h2)
	http.HandleFunc("/options", h3)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
