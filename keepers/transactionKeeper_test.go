package keepers

import (
	"testing"
	"time"
	types "tracker/types"
)

func TestWhenCreateTransactionKeeperStoreNotIsNil(t *testing.T) {
	tk := NewTransactionKeeper()
	if tk == nil {
		t.Error("pointer is nil")
	}

	if tk.store == nil {
		t.Error("store is nil, this is not valid")
	}
}

func TestWhenAddTransactionInKeeper(t *testing.T) {
	tk := NewTransactionKeeper()

	transaction := &types.Transaction{
		Id:    "qwgevjhvsabmvqwkjeh",
		Date:  time.Now(),
		Coin:  "Zillqa",
		Value: 15742121.53,
	}

	quantity := len(tk.store)
	tk.AddTransaction(transaction)
	newQuantity := len(tk.store)

	if newQuantity-quantity != 1 {
		t.Error("must add only one new transaction")
	}

	addTransaction := tk.store[transaction.Id]
	if addTransaction == nil {
		t.Error("transaction value is nil")
	}

}

func TestWhenCreateTransactionKeeperStoreNotIsNi(t *testing.T) {
	tk := NewTransactionKeeper()
	if tk == nil {
		t.Error("pointer is nil")
	}

	if tk.store == nil {
		t.Error("store is nil, this is not valid")
	}
}
