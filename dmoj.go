package main

import "time"

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

// type CodeforcesContest struct {
// 	Id                  int    `json:"id"`
// 	Name                string `json:"name"`
// 	Type                string `json:"type"`
// 	Phase               string `json:"phase"`
// 	Frozen              bool   `json:"fronzen"`
// 	DurationSeconds     int    `json:"durationSeconds"`
// 	StartTimeSeconds    int    `json:"startTimeSeconds"`
// 	RelativeTimeSeconds int    `json:"relativeTimeSeconds"`
// }

// func (c CodeforcesContest) Print() {
// 	fmt.Printf("Id: %v\n", c.Id)
// 	fmt.Printf("Name: %v\n", c.Name)
// 	fmt.Printf("Type: %v\n", c.Type)
// 	fmt.Printf("Phase: %v\n", c.Phase)
// 	fmt.Printf("Frozen: %v\n", c.Frozen)
// 	fmt.Printf("DurationSecond: %v\n", c.DurationSeconds)
// 	fmt.Printf("StartTimeSecond: %v\n", c.StartTimeSeconds)
// 	fmt.Printf("RelativeTimeSecond: %v\n", c.RelativeTimeSeconds)
// }

// func (cc CodeforcesContest) get_name() string {
// 	return cc.Name
// }

// func (cc CodeforcesContest) get_date() string {
// 	date := time.Now().Add(time.Duration(-1*cc.RelativeTimeSeconds) * time.Second)
// 	zone, _ := date.Zone()
// 	return fmt.Sprintf("%v %v, %v %v:%v %v", date.Month(), date.Day(), date.Year(), date.Hour(), date.Minute(), zone)
// }

// func (cc CodeforcesContest) get_url() string {
// 	return fmt.Sprintf("https://codeforces.com/contests/%v", cc.Id)
// }

// func (cc CodeforcesContest) get_seconds() int {
// 	return cc.RelativeTimeSeconds
// }

// func (cc CodeforcesContest) is_active() bool {
// 	return cc.Phase == "BEFORE"
// }

// func (cc CodeforcesContest) get_judge_name() string {
// 	return "Codeforces"
// }

// func to_contests(cc []CodeforcesContest) []Contest {
// 	var res []Contest
// 	for _, c := range cc {
// 		res = append(res, c)
// 	}
// 	return res
// }
