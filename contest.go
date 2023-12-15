package main

import "time"

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	} else {
// 		return a
// 	}
// }

func SetJudges() {
	JUDGES = append(JUDGES, "Codeforces")
	JUDGES = append(JUDGES, "Dmoj")
	JUDGES = append(JUDGES, "AtCoder")
	JUDGES = append(JUDGES, "Usaco")
}

func UpdateNeeded() bool {
	return time.Since(LAST_UPDATED).Hours() > 2
}

func UpdateContests() {
	CONTESTS = append(CONTESTS, filter(ToContests(GetCodeforces().Result), FilterIsUpcoming, nil)...)
	CONTESTS = append(CONTESTS, filter(DmojToContests(GetDmoj().Data.Objects), FilterIsUpcoming, nil)...)
	//CONTESTS = append(CONTESTS, filter(GetUsaco(), FilterIsUpcoming, nil)...)
	CONTESTS = append(CONTESTS, GetUsaco()...)
	CONTESTS = append(CONTESTS, GetAtCoder()...)
	LAST_UPDATED = time.Now()
}

type Contest interface {
	GetName() string
	GetDate() string
	GetUrl() string
	GetSeconds() int
	IsActive() bool
	GetJudgeName() string
}

type ByDate []Contest

func (a ByDate) Len() int {
	return len(a)
}

func (a ByDate) Less(i, j int) bool {
	return a[i].GetSeconds() < a[j].GetSeconds()
}

func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ByJudge []Contest

func (a ByJudge) Len() int {
	return len(a)
}

func (a ByJudge) Less(i, j int) bool {
	return a[i].GetJudgeName() < a[j].GetJudgeName()
}

func (a ByJudge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type Filter struct {
	condition func(Contest, []string) bool
}

var FilterIsUpcoming = Filter{func(c Contest, a []string) bool { return c.IsActive() }}
var FilterForJudge = Filter{
	func(c Contest, judges []string) bool {
		for _, j := range judges {
			if j == c.GetJudgeName() {
				return false
			}
		}
		return true
	},
}

func filter(contests []Contest, filter Filter, judges []string) []Contest {
	var filtered []Contest
	for _, c := range contests {
		if filter.condition(c, judges) {
			filtered = append(filtered, c)
		}
	}
	return filtered
}
