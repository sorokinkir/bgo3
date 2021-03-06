package card

import (
	"fmt"
)

// Transaction ....
type Transaction struct {
	ID       int64
	Amount   float64
	DateTime int64
	MCC      string
	Status   string
}

// Card ....
type Card struct {
	Transactions []*Transaction
}

// AddTransaction ....
func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, transaction)
	for i, row := range card.Transactions {
		fmt.Printf(
			"(ID в слайсе %v), ID=%v, сумма=%v, код MCC=%v, время=%v, статус=%v\n",
			i,
			row.ID,
			row.Amount,
			row.MCC,
			row.DateTime,
			row.Status,
		)
	}
}

// SumByMCC суммириует МСС коды, которые находятся в слайте
func SumByMCC(transactions []*Transaction, mcc []string) int64 {
	var result int64
	for _, value := range transactions {
		for _, code := range mcc {
			if value.MCC == code {
				result += int64(value.Amount)
			}
		}
	}

	return result
}
