package calendar

import (
	"contest_overview/contest"
	"fmt"
	"time"
)

type Week struct {
	Days []Day
}

type Day struct {
	Date     int
	Contests []contest.Contest
}

type Month struct {
	Name  string
	Weeks []Week
}

type Year struct {
	Name   int
	Months []Month
}

type CalendarContent struct {
	Years map[int]Year
}

type SendCalendar struct {
	MonthName    string
	Year         int
	MonthContent Month
}

func DisplayCalendar(content CalendarContent, year int, month time.Month) SendCalendar {
	var res SendCalendar
	res.MonthName = month.String()
	res.Year = year
	res.MonthContent = content.Years[year].Months[month-1]
	return res
}

func ContestsToCalendar(contests []contest.Contest) CalendarContent {
	currentYear := time.Now().Year()
	var calendar CalendarContent
	calendar.Years = make(map[int]Year)

	for y := currentYear; y <= currentYear+1; y++ {
		var year Year
		year.Name = y
		year.Months = make([]Month, 12)

		for m := 1; m <= 12; m++ {
			var month Month
			month.Name = time.Month(m).String()
			month.Weeks = make([]Week, 5) // Set the number of weeks to 5

			// Populate each week with 7 days
			for w := 0; w < 5; w++ {
				var week Week
				week.Days = make([]Day, 7)

				// Populate each day with the date and empty contests
				for d := 1; d <= 7; d++ {
					day := Day{Date: (w * 7) + d}
					day.Contests = make([]contest.Contest, 0)
					week.Days[d-1] = day
				}

				month.Weeks[w] = week
			}

			year.Months[m-1] = month
		}

		calendar.Years[y] = year
	}
	//for now just the starting date -> later Contest needs to implement GetEnding()
	for _, c := range contests {
		date, err := time.Parse("January 2, 2006 15:04 MST", c.GetDate())
		if err != nil {
			fmt.Println("parse calendar", err, c.GetJudgeName())
			continue
		}
		calendar.Years[date.Year()].Months[date.Month()-1].Weeks[(date.Day()-1)/7].Days[(date.Day()-1)%7].Contests = append(calendar.Years[date.Year()].Months[date.Month()-1].Weeks[(date.Day()-1)/7].Days[(date.Day()-1)%7].Contests, c)
	}
	//fmt.Println(calendar)
	return calendar
}
