package models

type Accounts struct{
	ID int `json:"-"`
	Owner string `json:"owner"`
	Balance int `json:"balance"`
	Currency string `json:"currency"`
}