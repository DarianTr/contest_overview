package contest

import (
	"encoding/json"
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
	var minute string
	var hour string
	if date.Minute() < 10 {
		minute = fmt.Sprintf("0%v", date.Minute())
	} else {
		minute = fmt.Sprintf("%v", date.Minute())
	}
	if date.Hour() < 10 {
		hour = fmt.Sprintf("0%v", date.Hour())
	} else {
		hour = fmt.Sprintf("%v", date.Hour())
	}
	return fmt.Sprintf("%v %v, %v %v:%v %v", date.Month(), date.Day(), date.Year(), hour, minute, zone)

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

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
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
