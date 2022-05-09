package handlers_test

import (
	"testing"
	"time"
	"tracker/handlers"
	"tracker/keepers"
	"tracker/types"
)

var tk *keepers.TransactionKeeper = keepers.NewTransactionKeeper()
var transactionHandler, _ = handlers.NewTransactionHandler(tk)

func TestCreateNewTransactionHandler(t *testing.T) {
	th, err := handlers.NewTransactionHandler(tk)

	if err != nil {
		t.Error(err)
	}

	if th == nil {
		t.Error("transaction handler cannot be nil")
	}
}

func TestCreateNewTransactionHandlerWithNilKeeper(t *testing.T) {
	_, err := handlers.NewTransactionHandler(nil)

	if err == nil {
		t.Error("must throw an error because keeper is nil")
	}
}

func TestTransactionHandlerGetAllTransactions(t *testing.T) {
	_, err := transactionHandler.Handler(types.MESSAGE_TRANSACTION_KEEPER_GET_ALL_TRANSACTIONS, nil)

	if err != nil {
		t.Error("must works, get all transactions require any")
	}
}

func TestTransactionHandlerGetTransactionWithExistanceId(t *testing.T) {
	id := "asdasdasdqwdasdqw"
	_, err := transactionHandler.Handler(types.MESSAGE_TRANSACTION_KEEPER_GET_TRANSACTION, id)

	if err != nil {
		t.Error("must works, exist this id by default")
	}
}

func TestTransactionHandlerGetTransactionWithNoExistanceId(t *testing.T) {
	id := "asdasdasdqwdasdqasdasdasdw"
	_, err := transactionHandler.Handler(types.MESSAGE_TRANSACTION_KEEPER_GET_TRANSACTION, id)

	if err == nil || err.Error() != "ID doesn't exist" {
		t.Error("should throw id not exists")
	}
}

func TestTransactionHandlerAddValidTransaction(t *testing.T) {
	transaction := &types.Transaction{
		Id:    "aszxcqwaxc",
		Date:  time.Now(),
		Coin:  "osmosis",
		Value: 5,
	}
	_, err := transactionHandler.Handler(types.MESSAGE_TRANSACTION_KEEPER_ADD_TRANSACTION, transaction)

	if err != nil {
		t.Error("must works, all parameters needed to add new transaction are valid")
	}
}

func TestTransactionHandlerRemoveTransaction(t *testing.T) {
	id := "asdasdasdqwdasdqw"
	_, err := transactionHandler.Handler(types.MESSAGE_TRANSACTION_KEEPER_REMOVE_TRANSACTION, id)

	if err != nil {
		t.Error("must works, all parameters needed to remove a transaction are valid")
	}
}

func TestTransactionHandlerUpdateTransaction(t *testing.T) {
	transaction := &types.Transaction{
		Id:    "asdasdasdqwdasdqw",
		Date:  time.Now(),
		Coin:  "OtroNombre",
		Value: 984324,
	}
	_, err := transactionHandler.Handler(types.MESSAGE_TRANSACTION_KEEPER_UPDATE_TRANSACTION, transaction)

	if err != nil {
		t.Error("must works, all parameters needed to update a transaction are valid")
	}
}

func TestTransactionHandlerWithInvalidMessage(t *testing.T) {
	_, err := transactionHandler.Handler("this message not exist", nil)

	if err != handlers.ErrInvalidMessage {
		t.Error("must throw an error, message is not valid")
	}
}
