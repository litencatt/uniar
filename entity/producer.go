package entity

type Producer struct {
	ID          int64
	ProviderID  int64
	IdentityId  string
	IsAdmin     bool
}

type ProducerScene struct {
	PhotographID int64
	MemberID     int64
	SsrPlus      bool
}
