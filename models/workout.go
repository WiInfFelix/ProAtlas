package models

type Workout struct {
	ID          int      `json:"id"`
	User        User     `json:"user"`
	Exercise    Exercise `json:"exercise"`
	Sets        int      `json:"sets"`
	Repetitions int      `json:"repetitions"`
	Weight      float32  `json:"weight"`
}
