package calendar

type Date struct {
	Number  int
	IsToday bool
}

type Week struct {
	Days []Date
}

type Calender struct {
	Year    int
	Month   string
	Numbers []int
}

var example = []Date{
	{Number: 1, IsToday: false},
	{Number: 2, IsToday: false},
	{Number: 3, IsToday: false},
	{Number: 4, IsToday: false},
	{Number: 5, IsToday: false},
	{Number: 6, IsToday: false},
	{Number: 7, IsToday: false},
	{Number: 8, IsToday: true},
	{Number: 9, IsToday: false},
	{Number: 10, IsToday: false},
}
