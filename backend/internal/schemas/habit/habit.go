package habit

type Input struct {
	Title    string `json:"title"`
	WeekDays []uint `json:"weekDays"`
}

type Output struct {
	ID string `json:"id"`
}
