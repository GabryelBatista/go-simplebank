package models

type Entries struct{
	ID int `json:"-"`
	Account_id int `json:"accountId"`
	Amount int `json:"amount"`
}