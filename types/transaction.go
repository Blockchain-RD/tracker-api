package types

import (
	"time"
)

type Transaction struct {
	Id    string
	Date  time.Time
	Coin  string
	Value float64
}
