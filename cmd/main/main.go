package main

import (
	"github.com/sorokinkir/bgo3/pkg/card"
	"time"
)

func main() {
	master := &card.Card{
		Transactions: []*card.Transaction{
			&card.Transaction{
				Id:       1,
				Amount:   735.55,
				DateTime: time.Now().Unix(),
				MCC:      "5411",
				Status:   "Операция в обработке",
			},
			&card.Transaction{
				Id:       2,
				Amount:   2_000.00,
				DateTime: time.Now().Unix(),
				MCC:      "0000",
				Status:   "Пополнения",
			},
			&card.Transaction{
				Id:       3,
				Amount:   1_203.91,
				DateTime: time.Now().Unix(),
				MCC:      "5812",
				Status:   "Операция в обработке",
			},
		},
	}

	transactionItem := card.Transaction{
		Id:       4,
		Amount:   50_000.00,
		DateTime: time.Now().Unix(),
		MCC:      "5533",
		Status:   "Обработано",
	}
	card.AddTransaction(master, &transactionItem)
}
