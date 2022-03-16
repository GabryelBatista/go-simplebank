package database

import (
	"context"

	"example.com/simplebank/config"
	"example.com/simplebank/models"
	"go.mongodb.org/mongo-driver/bson"
)

var transfers = client.Database(config.MongoDB().Database).Collection("transfers")

func GetTransfers(id int) ([]bson.M, error) {

	var results []bson.M
		if cursor, err := transfers.Find(context.TODO(), bson.M{"fromaccount": id}); err != nil {
			return nil, err
		}else{
			if err := cursor.All(context.TODO(), &results); err != nil {
				return nil, err
			}else{
				return results, nil
			}
			
		}
		
}

func PostTransfers(newTransfers *models.Transfers) error {

	if accountTransfers, err := GetByIdAccounts(newTransfers.FromAccount); err != nil {
		return err
	} else {

		if accountTransfers.Balance >= newTransfers.Amount {
			if err := PostEntries(&models.Entries{ID: newTransfers.ID, Account_id: newTransfers.FromAccount, Amount: -newTransfers.Amount}); err != nil {
				return err
			} else {
				if err := PostEntries(&models.Entries{ID: newTransfers.ID, Account_id: newTransfers.ToAccount, Amount: newTransfers.Amount}); err != nil {
					return err
				} else {
					if _, err := transfers.InsertOne(context.TODO(), newTransfers); err != nil {
						return err
					} else {
						return nil
					}
				}
			}
		} else {
			return err
		}
	}

}

func PutTransfers() {

}

func DeleteTransfers() {

}
