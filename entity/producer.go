package entity

type Producer struct {
	ID         int64
	ProducerID string
	Name       string
}

type ProducerScene struct {
	PhotographID int64
	MemberID     int64
	Have         int64
}
