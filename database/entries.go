package database

import (
	"context"

	"example.com/simplebank/config"
	"example.com/simplebank/models"
	"go.mongodb.org/mongo-driver/bson"
)


var entries = client.Database(config.MongoDB().Database).Collection("entries")

func GetEntries(id int) ([]bson.M, error)  {
	var result []bson.M
	if cursor, err := entries.Find(context.TODO(), bson.M{"account_id": id}) ; err != nil  {
		return nil, err
	}else{
		cursor.All(context.TODO(), &result)
		return result, nil
		
	}
	
	
}

func PostEntries(entrie *models.Entries) error  {
	
	
	if accountEntrie, err := GetByIdAccounts(entrie.Account_id); err != nil {
		return err
	}else{
		accountEntrie.Balance += entrie.Amount
		PutAccounts(&accountEntrie)
	if _, err := entries.InsertOne(context.TODO(), entrie); err != nil{
		return err
	}else{
		return nil
	}
	}
	
}

func PutEntries()  {
	
}

func DeleteEntries()  {
	
}