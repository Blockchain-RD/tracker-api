package handlers

import (
	"errors"
	"tracker/keepers"
	"tracker/types"
)

type TransactionHandler struct {
	keeper *keepers.TransactionKeeper
}

func NewTransactionHandler(k *keepers.TransactionKeeper) (*TransactionHandler, error) {
	if k == nil {
		return nil, errors.New("keeper cannot be nil")
	}

	return &TransactionHandler{
		keeper: k,
	}, nil
}

func (th *TransactionHandler) Handler(message string, data interface{}) (interface{}, error) {
	var result interface{}
	var err error
	switch message {
	case types.MESSAGE_TRANSACTION_KEEPER_GET_ALL_TRANSACTIONS:
		{
			result, err = th.keeper.GetAllTransaction()
		}
	case types.MESSAGE_TRANSACTION_KEEPER_GET_TRANSACTION:
		{
			id := data.(string)
			result, err = th.keeper.GetTransaction(id)
		}

	case types.MESSAGE_TRANSACTION_KEEPER_ADD_TRANSACTION:
		{
			newTransaction := data.(*types.Transaction)
			err = th.keeper.AddTransaction(newTransaction)
		}

	case types.MESSAGE_TRANSACTION_KEEPER_REMOVE_TRANSACTION:
		{
			id := data.(string)
			err = th.keeper.DeleteTransaction(id)
		}

	case types.MESSAGE_TRANSACTION_KEEPER_UPDATE_TRANSACTION:
		{
			editedTransaction := data.(*types.Transaction)
			err = th.keeper.UpdateTransaction(editedTransaction)
		}
	default:
		{
			return nil, errors.New("message is not valid")
		}
	}
	return result, err
}
