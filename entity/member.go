package entity

type Member struct {
	ID   int64
	Name string
}

type ProducerMember struct {
	ProducerID  int64
	MemberID    int64
	Name        string
	BondLevel   int64
	Discography int64
}
