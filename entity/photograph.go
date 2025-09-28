package entity

import "time"

type Photograph struct {
	ID           int64
	Name         string
	Abbreviation string
	PhotoType    string
	GroupID      int64
	ReleasedAt   time.Time
}

type PhotographWithDetails struct {
	ID           int64
	Name         string
	Abbreviation string
	PhotoType    string
	GroupID      int64
	ReleasedAt   time.Time
	NameForOrder string
	GroupName    string
}
