package keepers

import (
	"errors"
	types "tracker/types"
	"tracker/utils"
)

var (
	ErrPointerEmpty           = errors.New("pointer transaction is empty")
	ErrInvalidId              = errors.New("invalid ID")
	ErrSeekingId              = errors.New("error seeking value")
	ErrIdIsUsed               = errors.New("id is used")
	ErrStoreHaventCreated     = errors.New("store haven't created")
	ErrIdDoesntExists         = errors.New("ID doesn't exists")
	ErrReadingTransactionFile = errors.New("error reading transaction file")
	ErrSavingTransactionFile  = errors.New("error reading transaction file")
)

const PATH_TRANSACTION_FILE string = "./data/transactions.json"

type TransactionKeeper struct {
	store map[string]*types.Transaction
}

func NewTransactionKeeper() *TransactionKeeper {
	transactionKeeper := &TransactionKeeper{}
	transactionKeeper.store = make(map[string]*types.Transaction)

	store, err := utils.GetObjectFromJSONFile[map[string]*types.Transaction](PATH_TRANSACTION_FILE)
	if err != nil {
		panic(ErrReadingTransactionFile)
	}

	for _, v := range *store {
		transactionKeeper.store[v.Id] = &types.Transaction{
			Id:    v.Id,
			Date:  v.Date,
			Coin:  v.Coin,
			Value: v.Value,
		}
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
	tk.SaveFile()

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
	tk.SaveFile()

	return nil
}

func (tk *TransactionKeeper) UpdateTransaction(transaction *types.Transaction) error {
	result := tk.store[transaction.Id]
	if result == nil {
		return ErrIdDoesntExists
	}

	tk.store[transaction.Id] = transaction
	tk.SaveFile()

	return nil
}

func (tk *TransactionKeeper) SaveFile() {
	errStore := utils.SaveFile[map[string]*types.Transaction](PATH_TRANSACTION_FILE, tk.store)
	if errStore != nil {
		panic(ErrSavingTransactionFile)
	}
}
