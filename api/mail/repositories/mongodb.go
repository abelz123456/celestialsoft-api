package repositories

import (
	"context"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	Mongo *mongo.Database
	Log   log.Log
}

func (r *mongodb) Save(ctx context.Context, data entity.EmailSent) error {
	return nil
}

func (r *mongodb) GetOneByUID(ctx context.Context, uid string) (*entity.EmailSent, error) {
	return nil, nil
}

func (r *mongodb) GetCollection(ctx context.Context) ([]entity.EmailSent, error) {
	return nil, nil
}
