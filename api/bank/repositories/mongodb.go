package repositories

import (
	"context"
	"errors"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	Mongo *mongo.Database
	Log   log.Log
}

func (r *mongodb) GetCollection(ctx context.Context) ([]entity.Bank, error) {
	var results = make([]entity.Bank, 0)
	cursor, err := r.Mongo.Collection("bank").
		Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var result = new(entity.Bank)
		if err := cursor.Decode(&result); err != nil {
			r.Log.Error(err, "mongodb.GetCollection Exception", nil)
			cursor.Close(ctx)
			return nil, err
		}

		results = append(results, *result)
	}

	cursor.Close(ctx)
	return results, nil
}

func (r *mongodb) Create(ctx context.Context, data entity.Bank) (*entity.Bank, error) {
	collection := r.Mongo.Collection("bank")

	if _, err := collection.InsertOne(ctx, &data); err != nil {
		r.Log.Error(err, "mongodb.Create Exception", nil)
		return nil, err
	}

	return &data, nil
}

func (r *mongodb) GetOneByCode(ctx context.Context, code string) (*entity.Bank, error) {
	mongoResult := r.Mongo.Collection("bank").
		FindOne(ctx, bson.M{"bankCode": code})

	if err := mongoResult.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		r.Log.Error(err, "mongodb.GetOneByCode Exception", nil)
		return nil, err
	}

	var dataResult entity.Bank
	if err := mongoResult.Decode(&dataResult); err != nil {
		r.Log.Error(err, "mongodb.GetOneByCode Exception", nil)
		return nil, err
	}

	if dataResult.Oid == "" {
		return nil, nil
	}

	return &dataResult, nil
}

func (r *mongodb) GetOneByOid(ctx context.Context, oid string) (*entity.Bank, error) {
	mongoResult := r.Mongo.Collection("bank").
		FindOne(ctx, bson.M{"oid": oid})

	if err := mongoResult.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		r.Log.Error(err, "mongodb.GetOneByOid Exception", nil)
		return nil, err
	}

	var dataResult entity.Bank
	if err := mongoResult.Decode(&dataResult); err != nil {
		r.Log.Error(err, "mongodb.GetOneByOid Exception", nil)
		return nil, err
	}

	return &dataResult, nil
}

func (r *mongodb) UpdateOne(ctx context.Context, bank entity.Bank, newData entity.Bank) (*entity.Bank, error) {
	collection := r.Mongo.Collection("bank")

	updateData := bson.M{
		"$set": bson.M{
			"bankName": newData.BankName,
		},
	}

	_, err := collection.UpdateOne(ctx, bson.M{"oid": bank.Oid}, updateData)
	if err != nil {
		r.Log.Error(err, "mongodb.UpdateOne Exception", nil)
		return nil, err
	}

	return r.GetOneByOid(ctx, bank.Oid)
}

func (r *mongodb) Delete(ctx context.Context, bank entity.Bank) error {
	if _, err := r.Mongo.Collection("bank").DeleteOne(ctx, bson.M{"oid": bank.Oid}); err != nil {
		r.Log.Error(err, "mongodb.Delete Exception", nil)
		return err
	}

	return nil
}
