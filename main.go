package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

var client *http.Client
var domjAPIToken string

type Data struct {
	Contest []Contest
	Judges  []string
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

func main() {
	SetJudges()
	var _ Contest = AtCoderContest{}
	var _ Contest = DmojContest{}
	var _ Contest = CodeforcesContest{}
	SetDomjAPIToken()
	client = &http.Client{Timeout: 10 * time.Second}

	http.HandleFunc("/", h1)
	http.HandleFunc("/search", h2)
	http.HandleFunc("/options", h3)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
