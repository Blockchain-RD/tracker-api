package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"tracker/keepers"
	"tracker/types"
)

var (
	ErrCannotBeNil          = errors.New("keeper cannot be nil")
	ErrInvalidMessage       = errors.New("message is not valid")
	ErrMessageIsEmpty       = errors.New("message is empty")
	ErrConvertingJsonToType = errors.New("error Convirtiendo")
	ErrRequestingApi        = errors.New("error consultando")
	ErrReadingBody          = errors.New("error leyendo")
)

type TransactionHandler struct {
	keeper *keepers.TransactionKeeper
}

func NewTransactionHandler(k *keepers.TransactionKeeper) (*TransactionHandler, error) {
	if k == nil {
		return nil, ErrCannotBeNil
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
			return nil, ErrInvalidMessage
		}
	}
	return result, err
}

func (th *TransactionHandler) HandleFunc(rw http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	if message == "" {
		log.Fatalln(ErrMessageIsEmpty)
	}

	defer r.Body.Close()
	if r.Body == http.NoBody {
		executeHandle(message, th, rw, nil)
		return
	}

	s, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err, s)
	}
	var t interface{}

	if s != nil {
		err = json.Unmarshal(s, &t)
		if err != nil {
			log.Fatalln(err, s)
		}
	}
	nt := types.NewTransaction(t.(map[string]interface{}))
	executeHandle(message, th, rw, nt)
}

func executeHandle(message string, th *TransactionHandler, rw http.ResponseWriter, v interface{}) {
	r, err := th.Handler(message, v)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}

	rw.WriteHeader(200)
	rw.Header().Add("content-type", "application/json")
	rw.Write(response)
}
