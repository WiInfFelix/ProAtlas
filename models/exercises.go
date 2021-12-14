package models

type Exercise struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	TargetMuscle string `json:"target_muscle"`
}
