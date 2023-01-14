package entity

import "time"

type Photograph struct {
	Name         string
	Abbreviation string
	PhotoType    string
	GroupID      int
	ReleasedAt   time.Time
}
