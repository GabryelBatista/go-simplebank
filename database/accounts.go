package database

import (
	"context"

	"example.com/simplebank/config"
	"example.com/simplebank/models"
	"go.mongodb.org/mongo-driver/bson"
)

var accounts = client.Database(config.MongoDB().Database).Collection("accounts")

func GetAccounts() ([]bson.M, error) {
	var results []bson.M

	if cursor, err := accounts.Find(context.TODO(), bson.M{}); err != nil {

		return nil, err
	} else {
		cursor.All(context.TODO(), &results)
		return results, nil
	}

}

func GetByIdAccounts(id int) (models.Accounts, error) {
	var account models.Accounts
	
	if err := accounts.FindOne(context.TODO(), bson.M{"id": id}).Decode(&account);  err != nil {
		return account, err
	}else{
		return account, nil
	}
	
}

func PostAccounts(account *models.Accounts) error {

	if _, err := accounts.InsertOne(context.TODO(), account); err != nil {
		return err
	} else {
		return nil
	}

}

func PutAccounts(account *models.Accounts) error {

	if _, err := accounts.UpdateOne(context.TODO(), bson.M{"id": account.ID}, bson.M{"$set": account}); err != nil {
		return err
	} else {
		return nil
	}

}

func DeleteAccounts(id int) error {

	if _, err := accounts.DeleteOne(context.TODO(), bson.M{"id": id}); err != nil {
		return err
	} else {
		return nil
	}
}
