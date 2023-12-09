package main

import (
	"fmt"
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
	// PreparedBy         string `json:"preparedBy"`
	// WebsiteUrl         string `json:"websiteUrl"`
	// Description        string `json:"description"`
	// Difficulty         int    `json:"difficulty"`
	// Kind               string `json:"kind"`
	// IcpcRegion         string `json:"icpcRegion"`
	// Country            string `json:"country"`
	// City               string `json:"city"`
	// Season             string `json:"season"`
}

type ByDate []CodeforcesContest

func (a ByDate) Len() int {
	return len(a)
}

func (a ByDate) Less(i, j int) bool {
	return -1*a[i].RelativeTimeSeconds < -1*a[j].RelativeTimeSeconds
}

func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
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
