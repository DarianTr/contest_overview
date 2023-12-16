package calendar

import "contest_overview/contest"

type Date struct {
	Number  int
	IsToday bool
}

type CalendarContent struct {
	Number   int
	Contests []contest.Contest
}

type Week struct {
	Days []Date
}
