package entity

type Music struct {
	ID          int64
	Name        string
	Normal      int64
	Pro         int64
	Master      int64
	Length      int64
	ColorTypeID int64
	LiveID      int64
	MusicBonus  int64
}

type MusicWithDetails struct {
	ID          int64
	Name        string
	Normal      int64
	Pro         int64
	Master      int64
	Length      int64
	ColorTypeID int64
	LiveID      int64
	MusicBonus  int64
	LiveName    string
	ColorName   string
}