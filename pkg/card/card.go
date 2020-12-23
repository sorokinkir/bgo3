package card

type Transaction struct {
	Id       int64
	Amount   float64
	DateTime int64
	MCC      string
	Status   string
}

type Card struct {
	Transactions []*Transaction
}
