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

func (cc CodeforcesContest) get_name() string {
	return cc.Name
}

func (cc CodeforcesContest) get_date() string {
	date := time.Now().Add(time.Duration(-1*cc.RelativeTimeSeconds) * time.Second)
	zone, _ := date.Zone()
	return fmt.Sprintf("%v %v, %v %v:%v %v", date.Month(), date.Day(), date.Year(), date.Hour(), date.Minute(), zone)
}

func (cc CodeforcesContest) get_url() string {
	return fmt.Sprintf("https://codeforces.com/contests/%v", cc.Id)
}

func (cc CodeforcesContest) get_seconds() int {
	return cc.RelativeTimeSeconds
}

type Filter struct {
	condition func(CodeforcesContest) bool
}

func filter(contests []CodeforcesContest, filter Filter) []CodeforcesContest {
	var filtered []CodeforcesContest
	for _, c := range contests {
		if filter.condition(c) {
			filtered = append(filtered, c)
		}
	}
	return filtered
}
