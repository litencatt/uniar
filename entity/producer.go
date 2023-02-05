package entity

type Producer struct {
	ID          int64
	IdentityId  string
	DisplayName string
}

type ProducerScene struct {
	PhotographID int64
	MemberID     int64
	Have         int64
}
