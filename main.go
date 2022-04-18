package main

import (
	"log"
	"time"
	keepers "tracker/keepers"
	"tracker/types"
)

func main() {
	tk := keepers.NewTransactionKeeper()

	log.Println("---------------------------------------------")
	log.Println("All transactions")
	log.Println("---------------------------------------------")
	transactions, err := tk.GetAllTransaction()
	if err != nil {
		panic(err)
	}

	printTransactions(transactions)

	// ----------------------------------------------------------------------

	log.Println("---------------------------------------------")
	log.Println("Add transactions")
	log.Println("---------------------------------------------")
	t := &types.Transaction{
		Id:    "qwgevjhvsabmvqwkjeh",
		Date:  time.Now(),
		Coin:  "Zillqa",
		Value: 15742121.53,
	}
	err = tk.AddTransaction(t)
	if err != nil {
		panic(err)
	}
	transactions, _ = tk.GetAllTransaction()
	printTransactions(transactions)

	// ----------------------------------------------------------------------

	log.Println("---------------------------------------------")
	log.Println("Update transactions")
	log.Println("---------------------------------------------")
	t.Coin = "Band"
	t.Value = 55682
	err = tk.UpdateTransaction(t)
	if err != nil {
		panic(err)
	}

	transactions, err = tk.GetAllTransaction()
	if err != nil {
		panic(err)
	}
	printTransactions(transactions)

	// ----------------------------------------------------------------------

	log.Println("---------------------------------------------")
	log.Println("Delete transaction by ID")
	log.Println("---------------------------------------------")
	tk.DeleteTransaction(t.Id)
	transactions, err = tk.GetAllTransaction()
	if err != nil {
		panic(err)
	}
	printTransactions(transactions)
}

func printTransactions(t map[string]*types.Transaction) {
	for id, v := range t {
		log.Printf("%s -> Coin: %s, Value: %f ", id, v.Coin, v.Value)
	}
}
