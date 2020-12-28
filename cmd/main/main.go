package main

import (
	"fmt"
	"github.com/sorokinkir/bgo3/pkg/card"
	"time"
)

func main() {
	master := &card.Card{
		Transactions: []*card.Transaction{

			{
				ID:       1,
				Amount:   735.55,
				DateTime: time.Now().Unix(),
				MCC:      "5411",
				Status:   "Операция в обработке",
			},
			{
				ID:       2,
				Amount:   2_000.00,
				DateTime: time.Now().Unix(),
				MCC:      "0000",
				Status:   "Пополнения",
			},
			{
				ID:       3,
				Amount:   1_203.91,
				DateTime: time.Now().Unix(),
				MCC:      "5812",
				Status:   "Операция в обработке",
			},
		},
	}

	transactionItem := card.Transaction{
		ID:       4,
		Amount:   50_000.00,
		DateTime: time.Now().Unix(),
		MCC:      "5533",
		Status:   "Обработано",
	}

	fmt.Println("------ Пример использования AddTransaction ------")
	card.AddTransaction(master, &transactionItem)

	codesMCC := []string{
		"5533",
		"5812",
	}
	fmt.Println("------ Пример использования SumByMCC ------")
	total := card.SumByMCC(master.Transactions, codesMCC)
	fmt.Printf("Сумма составляет: %v руб.\n", total)

	category := card.TranslateMCC(master.Transactions[0].MCC)
	fmt.Printf("Категория: %v\n", category)

	fmt.Println("------ Пример использования LastNTransactions ------")
	lastTransactions := card.LastNTransactions(master, 3)
	fmt.Println(lastTransactions)
}
