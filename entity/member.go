package entity

type Member struct {
	ID   int64
	Name string
}

type ProducerMember struct {
	MemberID    int64
	Name        string
	BondLevel   int64
	Discography int64
}
