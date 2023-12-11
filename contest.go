package main

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

type Contest interface {
	get_name() string
	get_date() string
	get_url() string
	get_seconds() int
	is_active() bool
}

type ByDate []Contest

func (a ByDate) Len() int {
	return len(a)
}

func (a ByDate) Less(i, j int) bool {
	return abs(a[i].get_seconds()) < abs(a[j].get_seconds())
}

func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type Filter struct {
	condition func(Contest) bool
}

func filter(contests []Contest, filter Filter) []Contest {
	var filtered []Contest
	for _, c := range contests {
		if filter.condition(c) {
			filtered = append(filtered, c)
		}
	}
	return filtered
}
