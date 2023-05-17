package repositories

import (
	"context"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	Mongo *mongo.Client
	Log   log.Log
}

func (r *mongodb) Save(ctx context.Context, data entity.PermissionPolicyUser) (*entity.PermissionPolicyUser, error) {
	return nil, nil
}
