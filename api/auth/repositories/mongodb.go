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

func (r *mongodb) Save(ctx context.Context, data entity.PermissionPolicyUser) (*entity.PermissionPolicyUser, error) {
	collection := r.Mongo.Collection("permissionPolicyUser")

	if _, err := collection.InsertOne(ctx, &data); err != nil {
		r.Log.Error(err, "mongodb.Save Exception", nil)
		return nil, err
	}

	return &data, nil
}

func (r *mongodb) GetOneByEmail(ctx context.Context, email string) (*entity.PermissionPolicyUser, error) {
	mongoResult := r.Mongo.Collection("permissionPolicyUser").
		FindOne(ctx, bson.M{"emailname": email})

	if err := mongoResult.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		r.Log.Error(err, "mongodb.GetOneByEmail Exception", nil)
		return nil, err
	}

	var dataResult entity.PermissionPolicyUser
	if err := mongoResult.Decode(&dataResult); err != nil {
		r.Log.Error(err, "mongodb.GetOneByEmail Exception", nil)
		return nil, err
	}

	if dataResult.Oid == "" {
		return nil, nil
	}

	return &dataResult, nil
}
