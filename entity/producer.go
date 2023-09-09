package entity

type Producer struct {
	ID          int64
	ProviderID  int64
	IdentityId  string
}

type ProducerScene struct {
	PhotographID int64
	MemberID     int64
	Have         int64
}
