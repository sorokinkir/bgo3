package card

import "fmt"

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

func AddTransaction(card *Card, transaction *Transaction) {
	// TODO дописать добавление элементов в слайс
	fmt.Println(transaction)

	for i, row := range card.Transactions {
		fmt.Printf(
			"(ID в слайсе %v), ID=%v, сумма=%v, код MCC=%v, время=%v, статус=%v\n",
			i,
			row.Id,
			row.Amount,
			row.MCC,
			row.DateTime,
			row.Status,
		)
	}
}
