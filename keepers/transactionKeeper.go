package keepers

import (
	"errors"
	"time"
	types "tracker/types"
)

var (
	ErrPointerEmpty       = errors.New("pointer transaction is empty")
	ErrInvalidId          = errors.New("invalid ID")
	ErrSeekingId          = errors.New("error seeking value")
	ErrIdIsUsed           = errors.New("id is used")
	ErrStoreHaventCreated = errors.New("store haven't created")
	ErrIdDoesntExists     = errors.New("ID doesn't exists")
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
		return ErrPointerEmpty
	}

	if len(transaction.Id) <= 0 {
		return ErrInvalidId
	}

	val, err := tk.store[transaction.Id]
	if err && val == nil {
		return ErrSeekingId
	}

	if val != nil {
		return ErrIdIsUsed
	}

	tk.store[transaction.Id] = transaction
	return nil
}

func (tk *TransactionKeeper) GetAllTransaction() (map[string]*types.Transaction, error) {
	if tk.store == nil {
		return nil, ErrStoreHaventCreated
	}
	return tk.store, nil
}

func (tk *TransactionKeeper) GetTransaction(id string) (*types.Transaction, error) {
	result := tk.store[id]
	if result == nil {
		return nil, ErrIdDoesntExists
	}
	return result, nil
}

func (tk *TransactionKeeper) DeleteTransaction(id string) error {
	result := tk.store[id]
	if result == nil {
		return ErrIdDoesntExists
	}

	delete(tk.store, id)
	return nil
}

func (tk *TransactionKeeper) UpdateTransaction(transaction *types.Transaction) error {
	result := tk.store[transaction.Id]
	if result == nil {
		return ErrIdDoesntExists
	}

	tk.store[transaction.Id] = transaction
	return nil
}
