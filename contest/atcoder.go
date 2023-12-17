package contest

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type AtCoderContest struct {
	Name       string
	Duration   string
	StartTime  time.Time
	RatedRange string
	Url        string
}

func (a AtCoderContest) GetName() string {
	return a.Name
}

func (a AtCoderContest) GetDate() string {
	date := a.StartTime
	zone, _ := time.Now().Zone()
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

func (a AtCoderContest) GetUrl() string {
	return "https://atcoder.jp/" + a.Url
}

func (a AtCoderContest) GetSeconds() int {
	return int(time.Until(a.StartTime).Seconds())
}
func (a AtCoderContest) IsActive() bool {
	return true
}
func (a AtCoderContest) GetJudgeName() string {
	return "AtCoder"
}

func GetAtCoder() []Contest {
	var c = colly.NewCollector()
	var res []Contest
	const url = "https://atcoder.jp/contests"
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("#contest-table-upcoming table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			startTime := row.ChildText("td:nth-child(1)")
			contestName := row.ChildText("td:nth-child(2) a")
			duration := row.ChildText("td:nth-child(3)")
			// ratedRange := row.ChildText("td:nth-child(4)")
			contestURL := row.ChildAttr("td:nth-child(2) a", "href")
			time, _ := time.Parse("2006-01-02 15:04:05-0700", startTime)
			contest := AtCoderContest{
				Name:      contestName,
				Duration:  duration,
				StartTime: time,
				Url:       contestURL,
			}
			res = append(res, contest)
		})
	})
	c.Visit(url)
	return res
}
