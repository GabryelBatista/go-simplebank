package models

type Transfers struct{
	ID int `json:"-"`
	FromAccount int `json:"fromAccount"`
	ToAccount int `json:"toAccount"`
	Amount int `json:"amount"`
}