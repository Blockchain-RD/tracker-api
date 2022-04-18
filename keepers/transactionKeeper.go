package keepers

import (
	"errors"
	"time"
	types "tracker/types"
)

type TransactionKeeper struct {
	store map[string]*types.Transaction
}

func NewTransactionKeeper() *TransactionKeeper {
	transactionKeeper := &TransactionKeeper{}
	transactionKeeper.store = make(map[string]*types.Transaction)

	transactionKeeper.store["asdasdasdqwdasdqw"] = &types.Transaction{
		Id:    "asdasdasdqwdasdqw",
		Date:  time.Now(),
		Coin:  "osmosis",
		Value: 100000,
	}

	return transactionKeeper
}

func (tk *TransactionKeeper) AddTransaction(transaction *types.Transaction) error {
	if transaction == nil {
		return errors.New("pointer transaction is empty")
	}

	if len(transaction.Id) <= 0 {
		return errors.New("invalid ID")
	}

	val, err := tk.store[transaction.Id]
	if err && val == nil {
		return errors.New("error seeking value")
	}

	if val != nil {
		return errors.New("id is used")
	}

	tk.store[transaction.Id] = transaction
	return nil
}

func (tk *TransactionKeeper) GetAllTransaction() (map[string]*types.Transaction, error) {
	if tk.store == nil {
		return nil, errors.New("store haven't created")
	}
	return tk.store, nil
}

func (tk *TransactionKeeper) GetTransaction(id string) (*types.Transaction, error) {
	result := tk.store[id]
	if result == nil {
		return nil, errors.New("ID doesn't exist")
	}
	return result, nil
}

func (tk *TransactionKeeper) DeleteTransaction(id string) error {
	result := tk.store[id]
	if result == nil {
		return errors.New("ID doesn't exist")
	}

	delete(tk.store, id)
	return nil
}

func (tk *TransactionKeeper) UpdateTransaction(transaction *types.Transaction) error {
	result := tk.store[transaction.Id]
	if result == nil {
		return errors.New("ID doesn't exist")
	}

	tk.store[transaction.Id] = transaction
	return nil
}
