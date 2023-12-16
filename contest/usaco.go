package contest

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type UsacoContest struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
}

func (uc UsacoContest) GetDate() string {
	date := uc.StartTime
	zone, _ := date.Zone()
	var minute string
	if date.Minute() < 10 {
		minute = fmt.Sprintf("0%v", date.Minute())
	} else {
		minute = fmt.Sprintf("%v", date.Minute())
	}
	return fmt.Sprintf("%v %v, %v %v:%v %v", date.Month(), date.Day(), date.Year(), date.Hour(), minute, zone)
}

func (uc UsacoContest) GetName() string {
	return uc.Name
}

func (uc UsacoContest) GetSeconds() int {
	return int(time.Until(uc.StartTime).Seconds())
}

func (uc UsacoContest) IsActive() bool {
	return uc.StartTime.Compare(time.Now()) == 1
}

func (uc UsacoContest) GetUrl() string {
	return "http://www.usaco.org/index.php"
}

func (uc UsacoContest) GetJudgeName() string {
	return "Usaco"
}

func LineToContest(line string) Contest {
	split := strings.Split(line, ":")
	name := split[1]
	unformated_date := split[0]
	dateParts := strings.Split(unformated_date, "-")
	month := strings.Split(strings.Trim(dateParts[0], " "), " ")[0]
	var year string
	if month == "Dec" {
		year = "2023"
	} else {
		year = "2024"
	}
	layout := "Jan 2 2006"
	startingDate, err := time.Parse(layout, strings.Trim(dateParts[0]+" "+year, " "))
	if err != nil {
		fmt.Println("first", err)
		panic(nil)
	}
	end := fmt.Sprintf("%s %s %s", month, dateParts[1], year)
	endDate, err := time.Parse(layout, end)
	if err != nil {
		fmt.Println("second", err)
		panic(nil)
	}
	return UsacoContest{
		Name:      name,
		StartTime: startingDate,
		EndTime:   endDate,
	}
}

func GetUsaco() []Contest {
	var res []Contest
	url := "http://usaco.org/"
	c := colly.NewCollector()
	//need to make 2023-2024 with time.Year()
	c.OnHTML("div.panel", func(e *colly.HTMLElement) {
		if e.ChildText("h2") == "2023-2024 Schedule" {
			content := e.Text
			lines := strings.Split(content, "\n")
			res = append(res, LineToContest(lines[3]))
			res = append(res, LineToContest(lines[4]))
			res = append(res, LineToContest(lines[5]))
			res = append(res, LineToContest(lines[6]))
		}
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
