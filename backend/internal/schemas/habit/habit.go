package habit

type Input struct {
	Title    string `json:"title"`
	WeekDays []uint `json:"weekDays"`
}

type Output struct {
	ID        string     `json:"id"`
	Title     string     `json:"title,omitempty"`
	WeekDays  []uint     `json:"weekDays,omitempty"`
	DayHabits []DayHabit `json:"habits,omitempty"`
}

type DayHabit struct {
	ID      string
	DayID   string
	HabitID string
}
