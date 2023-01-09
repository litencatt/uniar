// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package repository

import (
	"database/sql"
	"time"
)

type CenterSkill struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

type ColorType struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

type Group struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

type Life struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

type Member struct {
	ID        int64
	Name      string
	FirstName sql.NullString
	GroupID   int64
	Phase     int64
	Graduated int64
	CreatedAt time.Time
}

type Music struct {
	ID          int64
	Name        string
	Normal      int64
	Pro         int64
	Master      int64
	Length      int64
	ColorTypeID int64
	LiveID      int64
	ProPlus     sql.NullInt64
	MusicBonus  sql.NullInt64
	SetlistID   sql.NullInt64
	CreatedAt   time.Time
}

type Photograph struct {
	ID           int64
	Name         string
	GroupID      int64
	Abbreviation string
	PhotoType    string
	ReleasedAt   interface{}
	CreatedAt    time.Time
}

type Producer struct {
	ID        int64
	CreatedAt time.Time
}

type ProducerMember struct {
	ProducerID              int64
	MemberID                int64
	BondLevelCurent         int64
	BondLevelCollectionMax  int64
	BondLevelSceneMax       int64
	DiscographyDiscTotal    int64
	DiscographyDiscTotalMax int64
	CreatedAt               time.Time
}

type ProducerOffice struct {
	ID          int64
	ProducerID  int64
	OfficeBonus sql.NullInt64
	CreatedAt   time.Time
}

type ProducerScene struct {
	ProducerID int64
	SceneID    int64
	Have       int64
	CreatedAt  time.Time
}

type Scene struct {
	ID             int64
	PhotographID   int64
	MemberID       int64
	ColorTypeID    int64
	VocalMax       int64
	DanceMax       int64
	PerformanceMax int64
	CenterSkill    sql.NullString
	ExpectedValue  sql.NullString
	SsrPlus        int64
	CreatedAt      time.Time
}

type Skill struct {
	ID                int64
	Name              string
	ComboUpPercent    sql.NullInt64
	DurationSec       sql.NullInt64
	ExpectedValue     float64
	IntervalSec       int64
	OccurrencePercent int64
	ScoreUpPercent    sql.NullInt64
	CreatedAt         time.Time
}
