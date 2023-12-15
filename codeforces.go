package main

import (
	"fmt"
	"time"
)

type CodeforcesResponse struct {
	Status string              `json:"status"`
	Result []CodeforcesContest `json:"result"`
}

type CodeforcesContest struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Type                string `json:"type"`
	Phase               string `json:"phase"`
	Frozen              bool   `json:"fronzen"`
	DurationSeconds     int    `json:"durationSeconds"`
	StartTimeSeconds    int    `json:"startTimeSeconds"`
	RelativeTimeSeconds int    `json:"relativeTimeSeconds"`
}

func (c CodeforcesContest) Print() {
	fmt.Printf("Id: %v\n", c.Id)
	fmt.Printf("Name: %v\n", c.Name)
	fmt.Printf("Type: %v\n", c.Type)
	fmt.Printf("Phase: %v\n", c.Phase)
	fmt.Printf("Frozen: %v\n", c.Frozen)
	fmt.Printf("DurationSecond: %v\n", c.DurationSeconds)
	fmt.Printf("StartTimeSecond: %v\n", c.StartTimeSeconds)
	fmt.Printf("RelativeTimeSecond: %v\n", c.RelativeTimeSeconds)
}

func (cc CodeforcesContest) GetName() string {
	return cc.Name
}

func (cc CodeforcesContest) GetDate() string {
	date := time.Now().Add(time.Duration(-1*cc.RelativeTimeSeconds) * time.Second)
	zone, _ := date.Zone()
	return fmt.Sprintf("%v %v, %v %v:%v %v", date.Month(), date.Day(), date.Year(), date.Hour(), date.Minute(), zone)
}

func (cc CodeforcesContest) GetUrl() string {
	return fmt.Sprintf("https://codeforces.com/contests/%v", cc.Id)
}

func (cc CodeforcesContest) GetSeconds() int {
	return -1 * cc.RelativeTimeSeconds
}

func (cc CodeforcesContest) IsActive() bool {
	return cc.Phase == "BEFORE"
}

func (cc CodeforcesContest) GetJudgeName() string {
	return "Codeforces"
}

func ToContests(cc []CodeforcesContest) []Contest {
	var res []Contest
	for _, c := range cc {
		res = append(res, c)
	}
	return res
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
