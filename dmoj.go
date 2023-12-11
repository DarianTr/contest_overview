package main

import (
	"fmt"
	"time"
)

type DmojResponse struct {
	ApiVersion string   `json:"api_version"`
	Method     string   `json:"method"`
	Fetched    string   `json:"fetched"`
	Data       DmojData `json:"data"`
}

type DmojData struct {
	CurrentObjectCount int           `json:"current_object_count"`
	ObjectsPerPage     int           `json:"objects_per_page"`
	PageIndex          int           `json:"page_index"`
	HasMmore           bool          `json:"has_more"`
	Objects            []DmojContest `json:"objects"`
	TotalObjects       int           `json:"total_objects"`
	TotalPages         int           `json:"total_pages"`
}

type DmojContest struct {
	Key       string    `json:"key"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	TimeLimit *float64  `json:"time_limit,omitempty"`
	IsRated   bool      `json:"is_rated"`
	RateAll   bool      `json:"rate_all"`
	Tags      []string  `json:"tags"`
}

var _ Contest = DmojContest{}

func (dc DmojContest) get_name() string {
	return dc.Name
}

func (dc DmojContest) get_date() string {
	date := dc.StartTime
	zone, _ := date.Zone()
	return fmt.Sprintf("%v %v, %v %v:%v %v", date.Month(), date.Day(), date.Year(), date.Hour(), date.Minute(), zone)
}

func (dc DmojContest) get_url() string {
	return fmt.Sprintf("https://dmoj.ca/contest/%v", dc.Key)
}

func (dc DmojContest) get_seconds() int {
	return int(time.Until(dc.StartTime).Seconds())
}

func (dc DmojContest) is_active() bool {
	return dc.EndTime.After(time.Now())
}

func (dc DmojContest) get_judge_name() string {
	return "Dmoj"
}

func dmoj_to_contests(dc []DmojContest) []Contest {
	var res []Contest
	for _, c := range dc {
		res = append(res, c)
	}
	return res
}
