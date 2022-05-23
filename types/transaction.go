package types

import (
	"errors"
	"fmt"
	"time"
)

type Transaction struct {
	Id    string
	Date  time.Time
	Coin  string
	Value float64
}

var (
	ErrIdDontExist    = errors.New("id don't exist")
	ErrCoinDontExist  = errors.New("coin don't exist")
	ErrDateDontExist  = errors.New("date don't exist")
	ErrConvertingDate = errors.New("date don't exist")
	ErrValueDontExist = errors.New("value don't exist")
)

func NewTransaction(m map[string]interface{}) *Transaction {
	t := &Transaction{}
	id, err := m["Id"]
	if !err {
		panic(ErrIdDontExist)
	}
	t.Id = fmt.Sprint(id)

	date, err := m["Date"]
	if !err {
		panic(ErrDateDontExist)
	}
	parseDate, e := time.Parse(time.RFC3339, fmt.Sprint(date))
	if e != nil {
		panic(ErrConvertingDate)
	}
	t.Date = parseDate

	coin, err := m["Coin"]
	if !err {
		panic(ErrCoinDontExist)
	}
	t.Coin = fmt.Sprint(coin)

	value, err := m["Value"]
	if !err {
		panic(ErrValueDontExist)
	}
	t.Value = value.(float64)
	return t
}
